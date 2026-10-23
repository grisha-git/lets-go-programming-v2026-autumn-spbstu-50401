package main

import "fmt"

func main() {
	var departmentCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Input reading error:", err)

		return
	}

}
