package main

func TwoSum(numbers []int, target int) [2]int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if j != i && j+i == target {
				break
			}
		}
	}

	return [2]int{i, j}
}
