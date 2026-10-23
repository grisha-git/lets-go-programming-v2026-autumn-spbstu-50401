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
			return
		}

		minimum := 15
		maximum := 30
		for j := range stuffNumber {
			var operator string
			var val int
			if _, err := fmt.Scanln(&operator, &val); err != nil {
				fmt.Println("Invalid input:", err)
				return
			}

			if operator != "<=" && operator != ">=" {
				fmt.Println("Invalid operator")
				return
			}
		}

	}
}
