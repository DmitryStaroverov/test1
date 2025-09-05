package main

import "fmt"

func main() {
	var num1, num2 float32
	var symbol string
	fmt.Scan(&num1, &num2, &symbol)

	switch symbol {
	case "+":
		fmt.Print(num1 + num2)
	case "-":
		fmt.Print(num1 - num2)
	case "*":
		fmt.Print(num1 * num2)
	case "/":

		if num2 == 0 {
			fmt.Print("Делить на ноль нельзя!")
		} else {
			fmt.Print(num1 / num2)
		}

	case "%":

		if num2 == 0 {
			fmt.Print("Делить на ноль нельзя!")
		} else {
			fmt.Print(float32(int(num1) % int(num2)))
		}

	}

}
