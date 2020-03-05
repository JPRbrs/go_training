package main

import (
	"fmt"
	"strings"
)

func accum(s string) string {
	parts := make([]string, len(s))

	for i := 0; i < len(s); i++ {
		parts[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
	}

	return strings.Join(parts, "-")
}

func main() {
	fmt.Println(accum("abcd"))    // should return "A-Bb-Ccc-Dddd"
	fmt.Println(accum("RqaEzty")) // should return "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
	fmt.Println(accum("cwAt"))    // should return "C-Ww-Aaa-Tttt"
}
