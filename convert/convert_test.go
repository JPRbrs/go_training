package convert

import (
	"testing"
)


func TestConvert(t *testing.T) {
	var tests = []struct {
		input string
		expected string
	} {
		{ "The-Stealth_Warrior", "TheStealthWarrior" },
		{ "the-Stealth-Warrior","theStealthWarrior" },
	}

	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			ans := ToCamelCase(tt.input)
			if ans != tt.expected{
				t.Errorf("Got: %s, expected %s", tt.input, tt.expected)
			}
		})
	}
}
