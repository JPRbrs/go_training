package main

import (
	"testing"
	"fmt"
)

func TestAlphanumericEmptyString(t *testing.T) {
	result := alphanumeric("")
	if result == true {
		t.Errorf("Empty string is not alphanumeric. Got %v expected False", result)
	}
}

func TestSingleLetter(t *testing.T) {
	result := alphanumeric("a")
	if result == false {
		t.Errorf("a is an alphanumeric string. Got %v expected true", result)
	}
}

func TestSingleDigit(t *testing.T) {
	result := alphanumeric("3")
	if result == false {
		t.Errorf("3 is an alphanumeric string. Got %v expected true", result)
	}
}


func TestAlphanumericWrongValues(t *testing.T){
	var tests = []struct {
		input string
		output bool
	}{
		{"asdfa_asdfas", true},
		{"asdfa-asdfas", true},
		{"asdfa/asdfas", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("String to test: %v", tt.input)
		t.Run(testname, func(t *testing.T) {
			result := alphanumeric(tt.input)
			if result != tt.output {
				t.Errorf("Got %v, expected %v", result, tt.output)
			}
		})
	}
}
