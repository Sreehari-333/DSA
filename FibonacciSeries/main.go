package main

import "fmt"

func fibonacciSeries() {

	a := 0
	b := 1
	fib := 0
	for i := 0; i <= 20; i++ {
		fmt.Println(a)
		fib = a + b
		a = b
		b = fib

	}
}

func main() {
	fibonacciSeries()
}
