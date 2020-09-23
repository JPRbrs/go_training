package main

import (
	"testing"
)

func TestCreatePhoneNumber(t *testing.T) {
	array := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	value := CreatePhoneNumber(array)

	if value != "(123) 456-7890" {
		t.Errorf("%v returned %s, wanted %s", array, value, "(123) 456-7890")
	} else {
		t.Log("All good")
	}
}

// func TestFunction(t *testing.T) {
// 	type tables struct {
// 		input    [10]uint
// 		expected string
// 	}
// 	var testTables = []tables{
// 		{{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, "(123) 456-7890"},
// 		{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, "(000) 000-0000"},
// 	}

// 	for _, table := range testTables {
// 		value := CreatePhoneNumber(table.input)
// 		if value != table.expected {
// 			t.Errorf("%v returned %s, wanted %s", table.input, value, table.expected)
// 		} else {

// 			t.Log("All good")
// 		}
// 	}
// }
