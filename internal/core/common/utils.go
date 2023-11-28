package common

import (
	"encoding/json"
	"strings"

	"github.com/volatiletech/null/v8"
)

func JSONOrDefault(j null.JSON) json.RawMessage {
	if !j.Valid || len(j.JSON) == 0 {
		return []byte(`{}`)
	}
	return j.JSON
}

func PrependFormulaTypeDefault(formula string) string {
	if len(formula) > 4 {
		if strings.HasPrefix(formula, DBCFormulaType.String()) {
			return formula
		}
		if strings.HasPrefix(formula, CustomFormulaType.String()) {
			return formula
		}
		return DBCFormulaType.String() + ": " + formula
	}
	return formula
}
