package main

import "fmt"

func main() {
	var number int
	var i int
	var j int
	var k int

	fmt.Println("How many number?")

	fmt.Scanln(&number)

	arr := make([]int, number)
	var evenSlice []int
	var oddSlice []int
	for i = 0; i < number; i++ {

		fmt.Println("Enter Number")
		fmt.Scanln(&arr[i])

	}
	for i = 0; i < len(arr); i++ {
		fmt.Println(arr[i])

		if arr[i]%2 == 0 {
			evenSlice = append(evenSlice, arr[i])
		} else {
			oddSlice = append(oddSlice, arr[i])
		}

	}

	fmt.Println("Even Number :")
	for j = 0; j < len(evenSlice); j++ {
		fmt.Println(evenSlice[j])
	}
	fmt.Println("odd Number :")
	for k = 0; k < len(oddSlice); k++ {
		fmt.Println(oddSlice[k])
	}

}
