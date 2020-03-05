package main

import (
	"fmt"
)

func dnaStrand(dna string) string {
	var DNAmap = map[rune]string{
		65: "T",
		84: "A",
		71: "C",
		67: "G",
	}
	var ret_val string
	for _, char := range dna {
		ret_val += DNAmap[char]
	}

	fmt.Println(ret_val)
	return ret_val
}

func main() {
	dnaStrand("ATTGC") // Should return TAACG
	dnaStrand("GTAT")  // Should return CATA
}
