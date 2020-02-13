package main

import "fmt"
import "math"

func findNb(m int) int {
	var cubeSideLength = 0
	for cubeSideLength = 1; m > 0; cubeSideLength++ {
		//fmt.Println("length", math.Pow(float64(cubeSideLength), 3))
		m -= int(math.Pow(float64(cubeSideLength), 3))
		//fmt.Println(m)
	}
	if m != 0 {
		cubeSideLength = -1
	}
	fmt.Println("cube length", cubeSideLength)
	return int(cubeSideLength)
}

func main() {
	fmt.Println(findNb(100)) // should return 5
}
