package DsaToolKit

import "fmt"

func insert(arr []int, val int, index int) []int {

	if index < 0 && index > len(arr) {
		fmt.Println("Invalid Index")
		return arr
	}

	arr = append(arr[:index], append([]int{val}, arr[index:]...)...)
	return arr
}

func delete(arr []int, index int) []int {

	if len(arr) == 0 {
		fmt.Println("Array is empty")
	}

	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

func search(arr []int, value int) int {

	if len(arr) < 0 {
		fmt.Println("Array is empty")
	}

	for i, val := range arr {

		if val == value {
			return i
		}
	}
	return 0
}

func Array() {

	arr := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("Orginal Array \n %d\n", arr)

	fmt.Println("After adding value")
	fmt.Println(insert(arr, 8, 7))

	fmt.Println("After Deleting value")
	fmt.Println(delete(arr, 6))

	fmt.Printf("Searched value is in %dth index\n", search(arr, 5))
	fmt.Println()
}
