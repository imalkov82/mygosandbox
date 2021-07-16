package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumOfSlice(a []int) int {
	sumResult := 0
	for _, v := range a {
		sumResult += v
	}
	return sumResult
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solution(a []int, k, l int) int {
	sumCache := make(map[string]int)
	var helper func(int, int, int, int) int
	helper = func(ks, ke, ls, le int) int {
		hash := strings.Join([]string{strconv.Itoa(ks), strconv.Itoa(ke), strconv.Itoa(ls), strconv.Itoa(le)}, "-")

		if _, ok := sumCache[hash]; !ok {
			sumCache[hash] = sumOfSlice(a[ks:ke]) + sumOfSlice(a[ls:le])
		}

		if ke == len(a) || le == len(a) {
			return sumCache[hash]
		}

		return max(sumCache[hash], helper(ks+1, ke+1, ls+1, le+1))
	}

	mmax := -1

	for i := 0; i < len(a)-k-l; i++ {
		mmax = max(mmax, helper(0, k, k+i, k+i+l))
	}

	return mmax
}

func main() {
	fmt.Println(solution([]int{6, 1, 4, 6, 3, 2, 7, 4}, 3, 2))
	fmt.Println(solution([]int{10, 19, 15}, 2, 2))
}
