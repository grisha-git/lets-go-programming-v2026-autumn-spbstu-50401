package main

import "fmt"

func main() {
	var departmentCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Input reading error:", err)

		return
	}

	for range departmentCount {
		var employeeCount int

		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Input reading error:", err)

			return
		}

	}
}
