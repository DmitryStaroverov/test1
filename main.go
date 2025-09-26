package main

import "fmt"

func main() {
	var arr [15]int
	var arrEven [15]int
	var arrOdd [15]int

	evenCount := 0
	oddCount := 0

	// Ввод чисел
	fmt.Println("Введите 15 чисел:")
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
		if arr[i]%2 == 0 {
			arrEven[evenCount] = arr[i]
			evenCount++
		} else {
			arrOdd[oddCount] = arr[i]
			oddCount++
		}
	}

	// Два счётчика для чётных и нечётных
	i, j := 0, 0

	// Вывод поочередно: чётный, нечётный
	for i < evenCount && j < oddCount {
		fmt.Print(arrEven[i], " ")
		i++
		fmt.Print(arrOdd[j], " ")
		j++
	}

	// Если остались чётные
	for i < evenCount {
		fmt.Print(arrEven[i], " ")
		i++
	}

	// Если остались нечётные
	for j < oddCount {
		fmt.Print(arrOdd[j], " ")
		j++
	}
}
