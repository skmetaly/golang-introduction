package main

import "fmt"
import "strconv"

func main() {
	/*
		bool

		string

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte // alias for uint8

		rune // alias for int32
		// represents a Unicode code point

		float32 float64

		complex64 complex128

	*/

	var age int = 40
	var favNum float64 = 1.6180339
	randNum := 1

	// Error if variables are not used
	fmt.Println(age, favNum, randNum)

	// Same as any other languages, float arithmetic operations are not precise

	var numberOne = 1.000
	var numberPoint = .9999

	fmt.Println(numberOne - numberPoint)

	// Strings are just "slices" of bytes. It doesn't hold 'characters'
	// That's why when looping through an array and try to print it's values
	// you'll get ugly output( they are neither ASCII nor UTF-8 )
	stringTest := "test"
	for i := 0; i < len(stringTest); i++ {
		fmt.Println("string [i"+strconv.Itoa(i)+"]", stringTest[i])
	}

	// Simplest way to loop through a string is using range loops
	// Go treats UTF-8 specially when used through range
	// It returns a rune
	for i, value := range stringTest {
		fmt.Println(value,string(value), i)
	}
}
