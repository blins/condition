package condition

const (
	OpTrue  = "true"
	OpFalse = "false"
	OpAnd   = "and"
	OpOr    = "or"
	OpNot   = "not"
)

// условие and
// возвращает true, если все низлежащие условия вернули true
type AndCondition struct {
	childs []Condition
}

func (condition *AndCondition) ParseArgs(args []string) ([]string, error) {
	c, a, err := ParseOne(args)
	condition.childs = append(condition.childs, c)
	return a, err
}

func (condition *AndCondition) Check(value interface{}) bool {
	for _, c := range condition.childs {
		if !c.Check(value) {
			return false
		}
	}
	return true
}

func (condition *AndCondition) AppendChild(childs ...Condition) {
	condition.childs = append(condition.childs, childs...)
}

func (condition *AndCondition) PrependChild(childs ...Condition) {
	condition.childs = append(childs, condition.childs...)
}

func AndConditionFabric() Condition {
	return &AndCondition{
		childs: make([]Condition, 0),
	}
}

// оператор or. Истино тогда, когда хотя бы одно из дочерних истинно
type OrCondition struct {
	childs []Condition
}

func (condition *OrCondition) ParseArgs(args []string) ([]string, error) {
	c, a, err := ParseOne(args)
	condition.childs = append(condition.childs, c)
	return a, err
}

func (condition *OrCondition) Check(value interface{}) bool {
	for _, c := range condition.childs {
		if c.Check(value) {
			return true
		}
	}
	return false
}

func (condition *OrCondition) AppendChild(childs ...Condition) {
	condition.childs = append(condition.childs, childs...)
}

func (condition *OrCondition) PrependChild(childs ...Condition) {
	condition.childs = append(childs, condition.childs...)
}

func OrConditionFabric() Condition {
	return &OrCondition{
		childs: make([]Condition, 0),
	}
}

// оператор not
type NotCondition struct {
	child Condition
}

func (condition *NotCondition) ParseArgs(args []string) ([]string, error) {
	c, a, err := ParseOne(args)
	condition.child = c
	return a, err
}

func (condition *NotCondition) Check(value interface{}) bool {
	return !condition.child.Check(value)
}

func (condition *NotCondition) AppendChild(childs ...Condition) {
	condition.child = childs[0]
}

func NotConditionFabric() Condition {
	return &NotCondition{}
}

// просто истина. просто заглушка.
type TrueCondition struct{}

func (condition *TrueCondition) ParseArgs(args []string) ([]string, error) {
	return args, nil
}

func (condition *TrueCondition) Check(value interface{}) bool {
	return true
}

func TrueConditionFabric() Condition {
	return &TrueCondition{}
}

// просто ложь. просто заглушка.
type FalseCondition struct{}

func (condition *FalseCondition) ParseArgs(args []string) ([]string, error) {
	return args, nil
}

func (condition *FalseCondition) Check(value interface{}) bool {
	return false
}

func FalseConditionFabric() Condition {
	return &FalseCondition{}
}

func init() {
	RegisterConditionFabric(OpAnd, FabricFunc(AndConditionFabric))
	RegisterConditionFabric(OpOr, FabricFunc(OrConditionFabric))
	RegisterConditionFabric(OpNot, FabricFunc(NotConditionFabric))
	RegisterConditionFabric(OpTrue, FabricFunc(TrueConditionFabric))
	RegisterConditionFabric(OpFalse, FabricFunc(TrueConditionFabric))
}
