package condition

import "testing"

type testCond struct {
	Args []string
	Data []interface{}
	Results []bool
	ALen int
	Err bool
}

var intValTestConds = []testCond {
	{
		Args:    []string{"eq", "5"},
		Data:    []interface{}{int(1), int(5), int(7), uint8(5), int64(45), int64(5), "Stop"},
		Results: []bool{false, true, false, true, false, true, false},
		ALen: 0,
		Err: false,
	},
	{
		Args:    []string{"5", "test"},
		Data:    []interface{}{int(1), int(5), int(7), uint8(5), int64(45), int64(5), "Stop"},
		Results: []bool{false, true, false, true, false, true, false},
		ALen: 1,
		Err: false,
	},
	{
		Args:    []string{"lt", "5", "test"},
		Data:    []interface{}{int(1), int(5), int(7), uint8(5), int64(45), int64(5), "Stop"},
		Results: []bool{true, false, false, false, false, false, false},
		ALen: 1,
		Err: false,
	},
	{
		Args:    []string{"lte", "5"},
		Data:    []interface{}{int(1), int(5), int(7), uint8(5), int64(45), int64(5), "Stop"},
		Results: []bool{true, true, false, true, false, true, false},
		ALen: 0,
		Err: false,
	},
	{
		Args:    []string{"lte", "rrr"},
		Data:    []interface{}{},
		Results: []bool{},
		ALen: 0,
		Err: true,
	},
}

func TestIntCondition_Check(t *testing.T) {
	for _, a := range intValTestConds {
		c := IntCondition{}
		_a, err := c.ParseArgs(a.Args)
		if (err != nil) != a.Err {
			t.Log(err)
			t.Fail()
		}
		if len(_a) != a.ALen {
			t.Fail()
		}
		for i, p := range a.Data {
			if c.Check(p) != a.Results[i] {
				t.Log(c)
				t.Logf("test %v, %t, %v", i, p, p)
				t.Fail()
			}
		}
	}
}