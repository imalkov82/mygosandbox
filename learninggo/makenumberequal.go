package main

import (
	"fmt"
	"math"
	"sort"
)

func sumWithoutIndex(a []int, pivotIdx int) int {
	sum := 0
	for i, el := range a {
		if i == pivotIdx {
			continue
		}
		sum += el * int(math.Abs(float64(pivotIdx-i)))
	}
	return sum
}

func solution(a []int) int {
	cnts := make([]int, 4)
	for _, i := range a {
		cnts[i-1] += 1
	}
	sort.Ints(a)
	return sumWithoutIndex(cnts, a[len(a)/2]-1)
}

func main() {
	fmt.Println(solution([]int{3, 2, 1, 1, 2, 3, 1}))
	fmt.Println(solution([]int{4, 1, 4, 1}))
	fmt.Println(solution([]int{3, 3, 3}))
}
