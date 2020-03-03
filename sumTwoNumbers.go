package main

import "fmt"

func twoSum(numbers []int, target int) [2]int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if j != i && j+i == target {
				fmt.Println("hola", i, j)
				return [2]int{i, j}
			}
		}
	}
	return [2]int{0, 0}
}

func main() {
	var numbers = []int{1, 2, 3, 4}
	var target = 4

	twoSum(numbers, target)
}
