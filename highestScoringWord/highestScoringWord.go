package main

import (
	"fmt"
	"strings"
)

func High(s string) string {
	words := strings.Split(s, " ")
	points_per_word := make([]int, len(words))

	for x := 0; x < len(words); x++ {
		word_total := 0
		for y := 0; y < len(words[x]); y++ {
			word_total += int(words[x][y]) - 96
		}
		points_per_word[x] = word_total

	}
	return words[highest_value_in_array(points_per_word)]
}

func highest_value_in_array(array []int) int {
	var max int
	var position int
	for pos := 0; pos < len(array); pos++ {
		if array[pos] > max {
			max = array[pos]
			position = pos
		}
	}
	return position
}

func main() {
	fmt.Println(High("this is a test"))
}
