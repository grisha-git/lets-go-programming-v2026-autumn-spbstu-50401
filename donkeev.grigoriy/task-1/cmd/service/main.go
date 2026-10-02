package main

import "fmt"

func main() {
	var firstOp int
	_, err1 := fmt.Scanln(&firstOp)
	if err1 != nil {
		fmt.Println("Invalid operand")
	}

	var secondOp int
	_, err2 := fmt.Scanln(&secondOp)
	if err2 != nil {
		fmt.Println("Invalid operand")
	}
}
