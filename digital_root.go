package main

import "fmt"

func digitalRoot(n int) int {
	if n < 10 {
		return n
	}
	var total = 0
	var resto = n % 10
	total += resto
	var modulo = n / 10
	for i := 0; modulo > 0; i++ {
		resto = modulo % 10
		modulo = modulo / 10
		total += resto
	}
	return digitalRoot(total)
}

func main() {
	fmt.Println(digitalRoot(30244))
}
