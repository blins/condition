/*
 */
package condition

/*
Проверка условий согласно правилам.

Задумка, что пользователь вводит правила на каком-то псевдо языке, по ним строится дерево или цепочка Condition,
которая потом проверяет на соответствие условиям
*/
import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNotExists    = errors.New("condition not exists")
	ErrArgsEmpty    = errors.New("args is empty")
	ErrRequiredLeft = errors.New("required left expression")
)

/*
Интерфейс, который должен реализовать интерпретатор условия
*/
type Condition interface {
	/*
		Функция забирает себе необходимые для условия параметры и возвращает остаток неразобранных аргументов для дальнейшей обработки.
		В случае возникновения ошибки, разбор выражения прекращается.
	*/
	ParseArgs(args []string) ([]string, error)
	/*
		Проверяет данные на соответствие условиям
	*/
	Check(value interface{}) bool
}

type RequredRight interface {
	/*
		Некоторые условия используют значения выражений справа. Эта функция вызывается, чтобы добавить поддерево в это условие
	*/
	AppendChild(childs ...Condition)
}

type RequredLeft interface {
	/*
		Некоторые условия используют значения выражений слева. Эта функция вызывается, чтобы добавить поддерево в это условие
	*/
	PrependChild(childs ...Condition)
}

/*
Фабрика условий. При разборе создаются условия по ключевым словам. Ключевое слово в Condition.ParseArgs не передаётся.
Поэтому об этом необходимо позаботиться, если это нужно
*/
type Fabric interface {
	Create() Condition
}

type FabricFunc func() Condition

func (fabric FabricFunc) Create() Condition {
	return fabric()
}

var (
	registryFabric map[string]Fabric
)

/*
зарегистрировать фабрику Condition
name должно быть уникально. Однако уникальность не проверяется. Это даёт возможность переопределить базовые условия
*/
func RegisterConditionFabric(name string, fabric Fabric) {
	if registryFabric == nil {
		registryFabric = make(map[string]Fabric)
	}
	name = strings.ToLower(name)
	registryFabric[name] = fabric
}

/*
Возвращает зарегистрированные фабрики
*/
func GetRegiteredConditionFabric() []string {
	res := make([]string, 0)
	for k := range registryFabric {
		res = append(res, k)
	}
	return res
}

// разобрать из псевдо языка одно Condition и ыернуть его и остаток нераспарсенных слов
func ParseOne(args []string) (Condition, []string, error) {
	if len(args) > 0 {
		if fabric, ok := registryFabric[args[0]]; ok {
			condition := fabric.Create()
			_args, err := condition.ParseArgs(args[1:])
			if err != nil {
				return nil, nil, err
			}
			return condition, _args, err
		} else {
			return nil, nil, ErrNotExists
		}
	}
	return nil, nil, ErrArgsEmpty
}

// расшифровать псевдо язык и вернуть цепочку или дерево условий
func ParseArgs(args []string) (Condition, error) {
	_args := args
	results := make([]Condition, 0)
	for len(_args) > 0 {
		condition, __args, err := ParseOne(_args)
		if err != nil {
			return nil, err
		}
		_args = __args
		if rl, ok := condition.(RequredLeft); ok {
			if len(results) == 0 {
				return nil, ErrRequiredLeft
			}
			rl.PrependChild(results[len(results)-1])
			results = results[0 : len(results)-1]
		}
		results = append(results, condition)
	}
	if len(results) == 1 {
		return results[0], nil
	}
	if len(results) == 0 {
		if fabric, ok := registryFabric[OpTrue]; ok {
			return fabric.Create(), nil
		} else {
			return nil, ErrNotExists
		}
	}
	if fabric, ok := registryFabric[OpAnd]; ok {
		condition := fabric.Create()
		if right, ok := condition.(RequredRight); ok {
			right.AppendChild(results...)
		}
		return condition, nil
	}
	return nil, ErrNotExists
}

type typeError string

func (e typeError) Error() string {
	return fmt.Sprintf("expected %v value", string(e))
}

func ErrExpectedValue(s string) error {
	return typeError(s)
}
