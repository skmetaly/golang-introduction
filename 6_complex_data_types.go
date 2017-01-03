package main

import "fmt"

func main() {

	// ARRAYS
	// Create an array of three ints.
	array := [3]int{10, 20, 30}

	// Loop over three ints and print them.
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	for i, value := range array {
		fmt.Println(i, value)
	}

	// Array of strings
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// SLICES
	simpleSlice := make([]int, 1, 2)
	// Appending after the initialized length values.
	// append will change the capacity if needed
	fmt.Println("Capacity:",cap(simpleSlice))
	simpleSlice = append(simpleSlice,2)
	simpleSlice = append(simpleSlice,2)
	fmt.Println("Capacity:",cap(simpleSlice))

	simpleSlice= append(simpleSlice,2)
	fmt.Println("SimpleSlice:", simpleSlice)

	numSlice := []int{1, 2, 3, 4, 5}
	for i, value := range numSlice {
		fmt.Println(i, value)
	}

	// Slice from slice
	pieceOfSlice := numSlice[2:5]
	for i := range pieceOfSlice {
		pieceOfSlice[i] = 150
	}

	for _, value := range pieceOfSlice {
		fmt.Println(value)
	}

	for _, value := range numSlice {
		fmt.Println(value)
	}

	// MAPS
	employeesAge := make(map[string]int)
	employeesAge["John"] = 43
	employeesAge["Bob"] = 22

	fmt.Println(len(employeesAge),employeesAge)

	/*
	 other types :
	 	function type
	 	interfaces
	 	struct
	 	channel
	 */

}
