package condition

import "reflect"

// разименовывает указатели в сами объекты. Т.е. позволяет дочерним условиям не задумываться о том, что они получают - ссылку или объект
type UnPtrCondition struct {
	next Condition
}

func (condition *UnPtrCondition) ParseArgs(args []string) ([]string, error) {
	return args, nil
}

func (condition *UnPtrCondition) Check(value interface{}) bool {
	v := reflect.Indirect(reflect.ValueOf(value))
	return condition.next.Check(v.Interface())
}

func UnPtrWrapCondition(cond Condition) Condition {
	return &UnPtrCondition{next: cond}
}
