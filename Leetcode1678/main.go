package main

import "fmt"

func interpret(command string) string {

	result := ""
	i := 0
	for i < len(command) {
		if string(command[i]) == "G" {
			result += "G"
			i++
		} else if string(command[i:i+2]) == "()" {
			result += "o"
			i += 2
		} else if string(command[i:i+4]) == "(al)" {
			result += "al"
			i += 4
		}

	}
	return result
}

func main() {

	command := "G()()()()(al)"
	fmt.Println(interpret(command))
}
