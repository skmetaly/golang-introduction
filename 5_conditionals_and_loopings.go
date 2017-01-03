package main

import "fmt"

func main() {

	// Both in conditionals and in for loops there is no need for parentheses around conditions
	// but curly braces are mandatory

	// FOR
	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// You can also `continue` to the next iteration of
	// the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}





	// IF
	// Basic example.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// You can have an `if` statement without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede conditionals; any variables
	// declared in this statement are available in all
	// branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}





	// SWITCH
	height := 5

	// Use switch on the height variable.
	// No default
	switch {
	case height <= 4:
		fmt.Println("Short")
	case height <= 5:
		fmt.Println("Normal")
	case height > 5:
		fmt.Println("Tall")
	}


	id := 10

	// Switch with multiple values in each case.
	switch id {
	case 10, 12, 14:
		fmt.Println("Even")
	case 11, 13, 15:
		fmt.Println("Odd")
	}

}
