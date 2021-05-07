package main

import (
	"testing"
	"strings"
)

func TestSplitStrings(t *testing.T) {
	value := SplitStrings("abc")
	expected := [2]string{"ab", "c_"}
	if (strings.Compare(value[0], expected[0]) == 1) || (strings.Compare(value[1], expected[1]) == 1) {
		t.Errorf("SplitString returned %v, expected %v", value, expected)
	}
}
