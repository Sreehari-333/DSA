package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) int {

	numOfStudents := 0

	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			numOfStudents++
		}
	}

	return numOfStudents
}

func main() {

	startTime := []int{1, 2, 3}
	endTime := []int{3, 2, 7}
	queryTime := 4

	fmt.Println(busyStudent(startTime, endTime, queryTime))
}
