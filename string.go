package condition

import "fmt"

/*
проверяет строку на соответствие функцией func (value string, arg string)

пример использования
 	condition.RegisterConditionFabric("prefix", condition.StringCheckerFabric(strings.HasPrefix))
	condition.RegisterConditionFabric("suffix", condition.StringCheckerFabric(strings.HasSuffix))
*/
type StringChecker struct {
	handler func(string, string) bool
	value   string
}

func (cond *StringChecker) ParseArgs(args []string) ([]string, error) {
	if len(args) == 0 {
		return nil, ErrExpectedValue("string")
	}
	cond.value = args[0]
	return args[1:], nil
}

func (cond *StringChecker) Check(value interface{}) bool {
	switch v := value.(type) {
	case string:
		return cond.handler(v, cond.value)
	case fmt.Stringer:
		return cond.handler(v.String(), cond.value)
	}
	return false
}

func StringCheckerFabric(handler func(string, string) bool) Fabric {
	return FabricFunc(func() Condition {
		return &StringChecker{handler: handler}
	})
}
