package main

import (
	"fmt"
)

func main() {
	fmt.Println(SplitStrings("abcdef"))
}

func SplitStrings (cadena string) []string {
	/*
	if len % 2 = 1 append_
	create slice for retval; lenght = len(cadena/2)
	for x in range(0, len(cadena), step 2):
	  retval.append(cadena[x] + cadena[x+1])
	return retval
	*/

	if (len(cadena) % 2 == 1) {
		cadena = cadena + string('_')
	}

	retval := make([]string, 0)
	for i := 0; i<len(cadena); i+=2 {
		retval = append(retval, string(cadena[i]) + string(cadena[i+1]))
	}
	return retval
}
