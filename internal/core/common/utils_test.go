package common

import "testing"

func TestPrependFormulaTypeDefault(t *testing.T) {

	tests := []struct {
		name    string
		formula string
		want    string
	}{
		{
			name:    "prepend if not present",
			formula: "31|8 dbc 321 blah",
			want:    "dbc: 31|8 dbc 321 blah",
		},
		{
			name:    "do not prepend if dbc",
			formula: "dbc: 31|8 dbc 321 blah",
			want:    "dbc: 31|8 dbc 321 blah",
		},
		{
			name:    "do not prepend if custom",
			formula: "custom: 31|8 dbc 321 blah",
			want:    "custom: 31|8 dbc 321 blah",
		},
		{
			name:    "do nothing if empty",
			formula: "",
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrependFormulaTypeDefault(tt.formula); got != tt.want {
				t.Errorf("PrependFormulaTypeDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
