package main

import (
	"fmt"
	"strings"
)

func accum(s string) string {
	var retVal string = ""
	for index, char := range s {
		for i := 0; i <= index; i++ {
			if i == 0 {
				retVal += strings.ToUpper(string(char))
			} else {
				retVal += strings.ToLower(string(char))
			}
		}
		retVal += "-"
	}
	return retVal[:len(retVal)-1]
}

func main() {
	fmt.Println(accum("abcd"))    // should return "A-Bb-Ccc-Dddd"
	fmt.Println(accum("RqaEzty")) // should return "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
	fmt.Println(accum("cwAt"))    // should return "C-Ww-Aaa-Tttt"
}
