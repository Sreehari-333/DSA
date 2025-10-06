package main

import (
	"fmt"
	"math/big"
)

func multiply(num1 string, num2 string) string {

	n1 := new(big.Int)
	n2 := new(big.Int)
	result := new(big.Int)

	n1.SetString(num1, 10)
	n2.SetString(num2, 10)

	result.Mul(n1, n2)

	return result.String()

}

func main() {

	num1 := "498828660196"
	num2 := "840477629533"
	fmt.Println(multiply(num1, num2))
}
