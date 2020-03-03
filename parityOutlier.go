package main

import "fmt"

func findOutlier(integers []int) int {
	odd := make([]int, 0)
	even := make([]int, 0)

	for _, number := range integers {
		if number%2 == 0 {
			even = append(even, number)
		} else {
			odd = append(odd, number)
		}
	}
	if len(odd) == 1 {
		return odd[0]
	}
	return even[0]
}

func main() {
	a := []int{2, 4, 0, 100, 4, 11, 2602, 36} // should return 11
	b := []int{160, 3, 1719, 19, 11, 13, -21} // should return 160
	fmt.Println(findOutlier(a))
	fmt.Println(findOutlier(b))
}
