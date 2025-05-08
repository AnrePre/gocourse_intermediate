// This program is made to practice and incorporate everything I have learnt in the intermediate section.
//Use closures, recursion, pointers, strings and runes, formatting verbs, fmt package, structs, methods, interfaces, struct embedding, generics

package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	//closures
	counter := closureFunc()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

}

func closureFunc() func() int {
	count := 0
	return func() int {
		count++
		return count
	}

}
