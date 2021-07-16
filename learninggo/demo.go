package main

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {
	// write your code in Go 1.4
	index := len(A) + 1
	res := make([]bool, len(A))
	for _, el := range A {
		if el > 0 && el < len(A) {
			res[el] = true
		}
	}
	for i := 1; i < len(A); i++ {
		if !res[i] {
			index = i
			break
		}
	}
	return index
}

func Solution2(A []int) int {
	sort.Ints(A)
	if len(A) == 0 || A[len(A)-1] < 0 {
		return 1
	}

	index := len(A) + 1

	for i := 0; i < len(A)-1; i++ {
		if A[i] < 0 || A[i+1] < 0 {
			continue
		}
		if A[i+1]-A[i] > 1 {
			index = A[i] + 1
			break
		}
	}
	return index
}

func main() {
	fmt.Println(Solution([]int{}), Solution2([]int{}))
	fmt.Println(Solution([]int{-1, -1, -1, 1, 1000}), Solution2([]int{-1, -1, -1, 1, 1000}))
	fmt.Println(Solution([]int{-1, -1, -1, -3}), Solution2([]int{-1, -1, -1, -3}))
}
