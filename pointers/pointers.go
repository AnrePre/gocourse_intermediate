//Pointers stores the memory address of another variable.
//referencing $ and dereferencing *
//allow us to modify the value of a variable indirectly
//var ptr *int (declaring the pointer)

package main

import "fmt"

func main() {
	
	var ptr *int
	var a int = 10
	ptr = &a


	fmt.Println(a)
	fmt.Println(ptr)
}
