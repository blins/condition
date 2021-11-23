package condition

import "testing"

func TestTimeCondition_Check(t *testing.T) {
	c := TimeCondition{}
	_, err := c.ParseArgs([]string{"2020-01-02T10:00", "10m"})
	if err != nil {
		t.Fatal(err)
	}
	tt, _ := anyParseTime("2020-01-01T00:00")
	if c.Check(tt) {
		t.Log(tt)
		t.Fail()
	}
	tt, _ = anyParseTime("2020-01-02T00:00")
	if c.Check(tt) {
		t.Log(tt)
		t.Fail()
	}
	tt, _ = anyParseTime("2020-01-02T10:00")
	if !c.Check(tt) {
		t.Log(tt)
		t.Fail()
	}
	tt, _ = anyParseTime("2020-01-02T10:03")
	if !c.Check(tt) {
		t.Log(tt)
		t.Fail()
	}
	tt, _ = anyParseTime("2020-01-02T10:10")
	if !c.Check(tt) {
		t.Log(tt)
		t.Fail()
	}
	tt, _ = anyParseTime("2020-01-02T10:10:01")
	if c.Check(tt) {
		t.Log(tt)
		t.Fail()
	}
}
