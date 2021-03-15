package main

import (
	"fmt"
	"strings"
)

func solve(str string) int {
	// var cadena string = "az"
	// fmt.Println(cadena[0]) // 97
	// fmt.Println(cadena[1]) // 122
	// restar 96
	// vocales 1, 5, 9, 15, 21

	var vowels = []int{1, 5, 9, 15, 21}
	for _, pos := range vowels {
		// What is cheaper: iterate vowels or iterate letters in word?
		// fmt.Println(pos + 96)
		// fmt.Println(string(pos + 96))
		str = strings.Replace(str, string(pos+96), "-", 3)
	}
	str = strings.Replace(str, "--", "-", 3)
	str = strings.Replace(str, "--", "-", 3)
	fmt.Println("len after replace")
	fmt.Println(len(str))
	fmt.Println(str)

	var splitString = strings.Split(str, "-")

	fmt.Println(splitString)
	fmt.Println(len(splitString))
	for _, char := range splitString {
		fmt.Println(char)
	}
	// parts := make([]string, len(splitString))

	// for _, group := range splitString {
	//	fmt.Println(group)
	//	fmt.Println(int(group - 96))
	// }

	return 1
}

func main() {
	solve("aeolian")  // should return 26
	solve("zodiac")   // should return 26
	solve("strength") // should return 26

	//fmt.Println(solve("zodiac")) // should return 26
	// fmt.Println(solve("strength")) // should return 57
}
