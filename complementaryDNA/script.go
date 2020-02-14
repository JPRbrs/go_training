package main

import (
	"fmt"
)

func dnaStrand(dna string) string {
	var DNAmap = map[string]string{
		"A": "T",
		"T": "A",
		"G": "C",
		"C": "G",
	}
	var ret_val string
	for _, char := range dna {
		ret_val += DNAmap[string(char)]
	}

	fmt.Println(ret_val)
	return ret_val
}

func main() {
	dnaStrand("ATTGC") // Should return TAACG
	dnaStrand("GTAT")  // Should return CATA
}
