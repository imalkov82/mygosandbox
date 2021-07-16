package main

import (
	"fmt"
)

func countOnes(n int) int {
	cnt := 0
	for {
		if n == 0 {
			break
		}
		cnt += n & 1
		n >>= 1
	}
	return cnt
}

func solution(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a == 1 {
		return countOnes(b)
	}
	if b == 1 {
		return countOnes(a)
	}

	return countOnes(a * b)
}
func main() {
	fmt.Println(solution(100000000, 100000000))
	fmt.Println(solution(100000000, 0))
}
