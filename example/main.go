package main

import (
	"fmt"
	"strings"

	"github.com/blins/condition"
)

type S struct {
	V int
	T string
}

func stringEquals(s1, s2 string) bool {
	return s1 == s2
}

type vFilter struct {
	child *condition.IntCondition
}

func (filter *vFilter) ParseArgs(args []string) ([]string, error) {
	return filter.child.ParseArgs(args)
}

func (filter *vFilter) Check(value interface{}) bool {
	switch v := value.(type) {
	case S:
		return filter.child.Check(v.V)
	}
	return false
}

func vFilterFabric() condition.Condition {
	return &vFilter{
		child: &condition.IntCondition{},
	}
}

type tFilter struct {
	child *condition.StringChecker
}

func (filter *tFilter) ParseArgs(args []string) ([]string, error) {
	if len(args) == 0 {
		return nil, condition.ErrArgsEmpty
	}
	shift := 1
	switch args[0] {
	case "prefix":
		filter.child.Handler = strings.HasPrefix
	case "suffix":
		filter.child.Handler = strings.HasSuffix
	case "equal":
		filter.child.Handler = stringEquals
	case "contains":
		filter.child.Handler = strings.Contains
	default:
		shift = 0
		filter.child.Handler = stringEquals
	}
	return filter.child.ParseArgs(args[shift:])
}

func (filter *tFilter) Check(value interface{}) bool {
	switch v := value.(type) {
	case S:
		return filter.child.Check(v.T)
	}
	return false
}

func tFilterFabric() condition.Condition {
	return &tFilter{
		child: &condition.StringChecker{},
	}
}

func init() {
	condition.RegisterConditionFabric("v", condition.FabricFunc(vFilterFabric))
	condition.RegisterConditionFabric("t", condition.FabricFunc(tFilterFabric))
}

func main() {
	cmdLine := strings.Fields("v > 5 and t prefix start")

	cond, _ := condition.ParseArgs(cmdLine)

	s1 := S{
		V: 3,
		T: "starts",
	}
	fmt.Printf("S: %v filter value is %v\n", s1, cond.Check(s1))
	s2 := S{
		V: 9,
		T: "starts",
	}
	fmt.Printf("S: %v filter value is %v\n", s2, cond.Check(s2))
	s3 := S{
		V: 9,
		T: "astarts",
	}
	fmt.Printf("S: %v filter value is %v\n", s3, cond.Check(s3))
}
