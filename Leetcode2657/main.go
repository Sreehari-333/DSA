package main

import "fmt"

func findThePrefixCommonArray(A []int, B []int) []int {

	freqA := make(map[int]int)
	freqB := make(map[int]int)
	result := []int{}

	for i := 0; i < len(A); i++ {

		freqA[A[i]]++
		freqB[B[i]]++

	}
	return result

}

func main() {

	A := []int{1, 3, 2, 4}
	B := []int{3, 1, 2, 4}
	fmt.Println(findThePrefixCommonArray(A, B))
}
