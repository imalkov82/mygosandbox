package main

import (
	"fmt"
	"math"
	"sort"
)

func solution(a []int) int {
	sm := 0
	sort.Ints(a)
	for i, el := range a {
		sm += int(math.Abs(float64(i + 1 - el)))
	}
	return sm
}

func main() {
	fmt.Println(solution([]int{2, 1, 4, 4}))
	fmt.Println(solution([]int{1, 2, 1}))
	fmt.Println(solution([]int{6, 2, 3, 5, 6, 3}))
	fmt.Println(solution([]int{6, 5, 4, 3, 2, 1}))
}
