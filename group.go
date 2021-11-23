package condition

import "fmt"

func parseGroupArgs(args []string, left string, right string) (Condition, []string, error) {
	_args := args
	rightIndex := 0
	stack := 0
	for i, a := range _args {
		if a == right && stack <= 0 {
			rightIndex = i
			break
		}
		if a == right {
			stack--
		}
		if a == left {
			stack++
		}
	}
	if rightIndex == 0 {
		return nil, nil, ErrExpected(right)
	}
	c, err := ParseArgs(_args[0:rightIndex])
	return c, _args[rightIndex+1:], err
}

// условие которое объединяет в группу одно условие. По факту вспомогательный парсер
type GroupCondition struct {
	left  string
	right string
	child Condition
}

func (condition *GroupCondition) ParseArgs(args []string) ([]string, error) {
	c, a, err := parseGroupArgs(args, condition.left, condition.right)
	condition.child = c
	return a, err
}

func (condition *GroupCondition) Check(value interface{}) bool {
	return condition.child.Check(value)
}

func (condition *GroupCondition) AppendChild(childs ...Condition) {
	condition.child = childs[0]
}

func GroupConditionFabric(left string, right string) (string, Fabric) {
	return left, FabricFunc(func() Condition {
		return &GroupCondition{
			left:  left,
			right: right,
		}
	})
}

func init() {
	RegisterConditionFabric(GroupConditionFabric("(", ")"))
	RegisterConditionFabric(GroupConditionFabric("{", "}"))
	RegisterConditionFabric(GroupConditionFabric("[", "]"))
}

type typeErrorExpected string

func (e typeErrorExpected) Error() string {
	return fmt.Sprintf("expected %v", string(e))
}

func ErrExpected(s string) error {
	return typeErrorExpected(s)
}
