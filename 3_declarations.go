package main

import "fmt"

func main() {

	// VARIABLES
	// full declaration
	var a string = "Hello world"
	fmt.Println("a = ", a)

	// multiple variables instantiated at once
	var f, e, intVal1, intVal2 int = 1, 2, 3, 4
	fmt.Println(f, e)

	// type inference
	// takes the data type based on initialisation
	// it also can't be changed later on. Once a boolean always a boolean
	var d = true
	fmt.Println(d)

	// Shorthand with type inference
	k := 100
	fmt.Println(k)

	// Multiple assignments
	f, e = intVal1, intVal2

	// Simple value inversions
	f, e = e, f

	var (
		g     = 2
		h int = 40
	)

	fmt.Println(g, h)

	// CONSTANTS
	// Typed
	const stringTypedConstant string = "Typed HelloWorld"

	// Untyped
	const stringConstant = "Hello, 世界"

	// After this declaration, hello is also an untyped string constant.
	// An untyped constant is just a value, one not yet given a defined type that would force it to obey the strict rules
	// that prevent combining differently typed values.
	fmt.Println(stringTypedConstant, "\n", stringConstant)

}
