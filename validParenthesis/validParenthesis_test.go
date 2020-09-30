package main

import "testing"

func TestValidParenthesis(t *testing.T) {
	type tables struct {
		input  string
		output bool
	}

	var testTables = []tables{
		{"()()()", true},
		{"((())", false},
		{"((())())", true},
		{"(((a)))", true},
		{"(((((()", false},
		{"(", false},
		{")", false},
		{"", true},
		{"())", false},
		{"()()()())", false},
		{"((((((((", false},
		{"())(", false},
		{")()()()(", false},
		{"(()()))(", false},
		{")()(", false},
		{"())(()", false},
	}
	for _, test := range testTables {
		value := ValidParenthesis(test.input)
		if value == test.output {
			t.Log("Test passed")
		} else {
			t.Errorf("ValidParenthesis returned %v, expected %v", value, test.output)
		}
	}
}
