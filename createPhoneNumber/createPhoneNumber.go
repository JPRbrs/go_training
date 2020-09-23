package main

import (
	"fmt"
)

func main() {
	array := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0} // returns "(123) 456-7890"}
	fmt.Println(CreatePhoneNumber(array))
}

func CreatePhoneNumber(numbers [10]uint) (phone string) {
	// Uses named return
	// Uses Sprintf to format string
	phone = fmt.Sprintf("(%s) %s-%s",
		translate(numbers[0:3]),
		translate(numbers[3:6]),
		translate(numbers[6:10]))
	return
}

func translate(slice []uint) string {
	// Uses magic number 48 to convert uint to a letter
	retval := ""
	for _, v := range slice {
		retval += string(v + 48)
	}
	return retval
}
