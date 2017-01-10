package main

import "fmt"
import (
	s "strings"
	"math"
	"time"
)

func main(){
	var p = fmt.Println

	// Strings
	// Here's a sample of the functions available in
	// `strings`. Since these are functions from the
	// package, not methods on the string object itself,
	// we need pass the string in question as the first
	// argument to the function. You can find more
	// functions in the [`strings`](http://golang.org/pkg/strings/)
	// package docs.
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))

	// Not part of `strings`, but worth mentioning here, are
	// the mechanisms for getting the length of a string in
	// bytes and getting a byte by index.
	p("Char:", "hello"[1])
	p("Len: ", len("hello"))


	// Math
	negative := -10
	fmt.Println(negative)
	// Use math.Abs to convert to a positive number.
	// ... We first convert to a float64.
	result := math.Abs(float64(negative))
	fmt.Println(result)

	resultPow := math.Pow(2, 3)
	fmt.Println(resultPow)

	value1 := 1.23
	fmt.Println(value1)
	fmt.Println(math.Floor(value1))

	value2 := 2.99
	fmt.Println(value2)
	fmt.Println(math.Floor(value2))

	value3 := -1.1
	fmt.Println(value3)
	fmt.Println(math.Floor(value3))

	fmt.Println(time.Now())
}
