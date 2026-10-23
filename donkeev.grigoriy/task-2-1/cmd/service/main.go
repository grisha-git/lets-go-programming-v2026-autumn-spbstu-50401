package main

import "fmt"

func main() {
	var departmentNumber int
	if _, err := fmt.Scanln(&departmentNumber); err != nil {
		fmt.Println("Invalid N:", err)
		return
	}

	for i := range departmentNumber {
		var stuffNumber int
		if _, err := fmt.Scanln(&stuffNumber); err != nil {
			fmt.Println("Invalid stuff number:", err)
		}

	}
}
