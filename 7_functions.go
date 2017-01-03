package main

import "fmt"

func main() {

	// simple function
	listNums := []int{1, 2, 3, 4, 5}
	sum := addThemUp(listNums)
	fmt.Println("Sum:", sum)

	// multiple returns
	num1, num2, num31 := next3Vals(5)
	fmt.Println(num1, num2, num31)

	// variadic functions
	fmt.Println("Sub:", subtractThem(1, 2, 3, 4, 5))

	//closures
	num3 := 3

	doubleNum := func() int {
		num3 *= 2
		return num3
	}

	fmt.Println("Double", doubleNum())
	fmt.Println("Double", doubleNum())

	// Deferring a function
	defer printString("i am last")
}

func addThemUp(nums []int) int {
	sum := 0

	for _, value := range nums {
		sum += value
	}

	return sum
}

func next3Vals(num int) (int, int, int) {
	return num + 1, num + 2, num + 3
}

func subtractThem(args ... int) int {

	finalValue := 0
	for _, value := range args {
		finalValue -= value
	}

	return finalValue
}

func printString(str string) {
	fmt.Println(str)
}
