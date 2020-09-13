package main

import (
	"fmt"
	"strings"
)

func wordScore(s string) (scor byte) {
	for i := range s {
		scor += s[i] - 96
	}
	return
}

func High(s string) string {
	var scor, scorNew byte
	var word string
	for _, wd := range strings.Split(s, " ") {
		scorNew = wordScore(wd)
		if scorNew > scor {
			scor = scorNew
			word = wd
		}
	}
	return word
}

func main() {
	fmt.Println(High("This is just a test"))
}
