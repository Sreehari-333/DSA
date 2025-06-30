package main

import "fmt"

func main() {

	arr := []int{1, 1, 2, 2, 3, 4, 5, 6, 7, 7, 8, 8, 9, 9}
	var dupli []int

	left := 0

	for left < len(arr) {
		if left+1 < len(arr) {
			if arr[left] != arr[left+1] {
				dupli = append(dupli, arr[left])
			}
		} else {
			dupli = append(dupli, arr[left])
		}
		left++
	}

	fmt.Println(dupli)
}
