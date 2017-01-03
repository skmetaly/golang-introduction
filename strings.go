package main

import (
	"fmt"
	"strings"
)
func main() {

	sampleStr := "Hello world"

	fmt.Println(strings.Contains(sampleStr,"Hello"))
	fmt.Println(strings.Index(sampleStr,"w"))
	fmt.Println(strings.Count(sampleStr,"o"))
	fmt.Println(strings.Replace(sampleStr,"world","jadu",1))


	// Similar to explode/implode
	sampleStrList := "1,2,3,4,5"
	fmt.Println(strings.Split(sampleStrList,","))

	fmt.Println(strings.Join([]string{"3","2","1"},","))


}
