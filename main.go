package main

import "fmt"

var a int
var b int

func main() {
	var arr [15]int
	var arr1 [15]int
	var arr2 [15]int

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])

		if arr[i]%2 == 0 {
			arr2[a] = arr[i]
			a++
		} else if arr[i]%2 != 0 {
			arr1[b] = arr[i]
			b++
		}

	}
	for i := 0; i < len(arr); i++ {
		if arr2[i] != 0 {
			fmt.Print(arr2[i], " ")
		}
		if arr1[i] != 0 {
			fmt.Print(arr1[i], " ")

		}

	}
}
