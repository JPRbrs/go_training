package main

import "fmt"

func MultiplicationTable(size int) [][]int {
	table := make([][]int, 0, size)
	for row := 1; row <= size; row++ {
		var array = make([]int, 0, 3)
		for y := 1; y <= size; y++ {
			array = append(array, y*row)
		}
		table = append(table, array)

	}
	return table

}

func main() {
	fmt.Println(MultiplicationTable(1))
}
