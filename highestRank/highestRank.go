package main

import "fmt"

func HighestRank(nums []int) int {
	frequencies := make(map[int]int)
	for _, value := range nums {
		if frequencies[value] == 0 {
			frequencies[value] = 1
		} else {
			frequencies[value]++
		}
	}
	fmt.Println(frequencies)

	frequency := 0
	higher := 0
	for key, value := range frequencies {
		if value > frequency {
			frequency = value
			higher = key
		}
	}
	return higher
}

func main() {
	a := []int{12, 10, 8, 12, 7, 6, 4, 10, 12, 10}
	fmt.Println(HighestRank(a))
}
