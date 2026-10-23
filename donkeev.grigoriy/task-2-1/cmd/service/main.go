package main

import "fmt"

func main() {
	var departmentNumber int
	if _, err := fmt.Scanln(&departmentNumber); err != nil {
		fmt.Println("Invalid N:", err)
		return
	}

	for range departmentNumber {
		var stuffNumber int
		if _, err := fmt.Scanln(&stuffNumber); err != nil {
			fmt.Println("Invalid stuff number:", err)
			return
		}

		minimum := 15
		maximum := 30
		incompatible := false
		for range stuffNumber {
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

			if incompatible {
				fmt.Println(-1)
				continue
			}
			if operator == ">=" {
				if val > maximum {
					fmt.Println(-1)
					incompatible = true
					continue
				}
				minimum = max(val, minimum)
			}
			if operator == "<=" {
				if val < minimum {
					fmt.Println(-1)
					incompatible = true
					continue
				}
				maximum = min(val, maximum)
			}

			fmt.Println(minimum)
		}

	}
}
