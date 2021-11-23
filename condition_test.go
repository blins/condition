package condition

import (
	"testing"
)

func intFabric() Condition {
	return &IntCondition{}
}

func TestParseArgs(t *testing.T) {
	RegisterConditionFabric("int", FabricFunc(intFabric))

	args := []string{"int", "lte", "5"}
	c, err := ParseArgs(args)
	if err != nil {
		t.Fatal(err)
	}
	if c == nil {
		t.Fatal()
	}
	if c.Check(int(5)) != true {
		t.Fail()
	}
}

func TestParseArgs2(t *testing.T) {
	RegisterConditionFabric("int", FabricFunc(intFabric))

	args := []string{"int", "lte", "5" , "and", "int", "gt", "3"}
	c, err := ParseArgs(args)
	if err != nil {
		t.Fatal(err)
	}
	if c == nil {
		t.Fatal()
	}
	if c.Check(int(5)) != true {
		t.Fail()
	}
	if c.Check(int(3)) != false {
		t.Fail()
	}
	if c.Check(int(4)) != true {
		t.Fail()
	}
}

func TestParseArgs3(t *testing.T) {
	RegisterConditionFabric("int", FabricFunc(intFabric))

	args := []string{"(", "int", "5", "or", "int", "18", ")" , "and", "int", "gt", "10"}
	c, err := ParseArgs(args)
	if err != nil {
		t.Fatal(err)
	}
	if c == nil {
		t.Fatal()
	}
	if c.Check(int(15)) != false {
		t.Fail()
	}
	if c.Check(int(13)) != false {
		t.Fail()
	}
	if c.Check(int(18)) != true {
		t.Fail()
	}
	if c.Check(int(17)) != false {
		t.Fail()
	}
	if c.Check(int(11)) != false {
		t.Fail()
	}
}
