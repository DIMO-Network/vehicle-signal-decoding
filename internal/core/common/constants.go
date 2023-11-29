package common

type FormulaType string

const (
	DBCFormulaType    FormulaType = "dbc"
	CustomFormulaType FormulaType = "custom"
)

func (t FormulaType) String() string {
	return string(t)
}
