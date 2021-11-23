package condition

import (
	"strconv"
)

const (
	opNoop = iota
	opLt
	opLte
	opEq
	opGte
	opGt
)

var (
	tableOp = map[string]int{
		"eq": opEq,
		"=":  opEq,
		"==": opEq,

		"lt": opLt,
		"<":  opLt,

		"lte": opLte,
		"<=":  opLte,
		"=<":  opLte,

		"gt": opGt,
		">":  opGt,

		"gte": opGte,
		">=":  opGte,
		"=>":  opGte,
	}
)

func opFromTable(op string) int {
	if o, ok := tableOp[op]; ok {
		return o
	}
	return opNoop
}

// сравнение числа
// Не является самостоятельным условием и используется для построения более сложных конструкций с определенным типом
type IntCondition struct {
	op    int
	value int
}

func (condition *IntCondition) ParseArgs(args []string) ([]string, error) {
	if len(args) == 0 {
		return nil, ErrExpectedValue("int")
	}
	i, err := strconv.Atoi(args[0])
	shift := 1
	if err != nil {
		if len(args) > 1 {
			i, err = strconv.Atoi(args[1])
		}
		if err != nil {
			return nil, ErrExpectedValue("int")
		}
		shift++
		condition.op = opFromTable(args[0])
	}
	if condition.op == opNoop {
		condition.op = opEq
	}
	condition.value = i
	return args[shift:], nil
}

func (condition *IntCondition) Check(value interface{}) bool {
	var i int
	switch v := value.(type) {
	case int:
		i = v
	case uint:
		i = int(v)
	case uint8:
		i = int(v)
	case int8:
		i = int(v)
	case uint16:
		i = int(v)
	case int16:
		i = int(v)
	case uint32:
		i = int(v)
	case int32:
		i = int(v)
	case uint64:
		i = int(v)
	case int64:
		i = int(v)
	default:
		return false
	}
	switch condition.op {
	case opEq:
		return i == condition.value
	case opLt:
		return i < condition.value
	case opLte:
		return i <= condition.value
	case opGt:
		return i > condition.value
	case opGte:
		return i >= condition.value
	}
	return false
}

/*
фабрика для самостоятельного использования IntCondition

Пример использования:
	condition.RegisterConditionFabric("gt", condition.IntFabric("gt"))
	condition.RegisterConditionFabric("gte", condition.IntFabric("gte"))
	condition.RegisterConditionFabric("lt", condition.IntFabric("lt"))
	condition.RegisterConditionFabric("lte", condition.IntFabric("lte"))
	condition.RegisterConditionFabric("eq", condition.IntFabric("eq"))

*/
func IntFabric(op string) Fabric {
	return FabricFunc(func() Condition {
		return &IntCondition{op: opFromTable(op)}
	})
}
