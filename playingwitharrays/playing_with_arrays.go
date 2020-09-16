package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

func main() {
	fmt.Println("sortScores result:", sortScores())

}

func arrayTest() {
	array := [...]int{1, 2, 3}
	scores := []int{1, 2, 3, 4, 5}
	slice := scores[2:4]
	fmt.Println(reflect.TypeOf(scores))
	fmt.Println(reflect.TypeOf(array))
	fmt.Println("slice:", slice)
	slice[0] = 999
	fmt.Println("scores:", scores)
}

func sortScores() []int {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)
	worst := make([]int, 5)
	fmt.Println("worse:", worst[:5])
	copy(worst[2:4], scores[:5])
	fmt.Println("scores:", scores[2:5])
	return worst
}
