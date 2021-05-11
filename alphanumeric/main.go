// _ and - were not allowed, it was easier that that

package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(alphanumeric("test"))
}

func alphanumeric (str string) bool {
	if str == "" {
		return false
	}
	for _, c := range str {
		if (!unicode.IsDigit(c) && !unicode.IsLetter(c) && !(c == '-') && !(c == '_')) {
			return false
		}
	}
	return true
}

// func alphanumeric(str string) bool {
//   return regexp.MustCompile("^[[:alnum:]]+$").MatchString(str)
// }
