//Create a Program that is run in the terminal. Where the user is taught the whole intermediate section of the go course. Thus add explanations and also examples of the code. This is a cool idea because when I get rusty on a section then I can just run the section that I would like to revise and then I will get the answers.

package main

import "fmt"

func main() {
	fmt.Println(".\n.\n.\nWelcome to My Golang Revision Tool.")
	fmt.Println("Last updated 05/09/2025")
	//Intro
loop:
	for {
		fmt.Println("\n---\nPlease select the chapter that you would like to revisit")
		fmt.Print("1. Closures\n2. Recursion\n3. Pointers\nYour answer: ")
		var answer int
		fmt.Scan(&answer)

		//Switch over choices
		switch answer {
		case 1:
			//closures
			fmt.Println("---\nClosures")
			fmt.Println("Definition: A closure in Go is a function that captures and uses variables from its surrounding scope, even after that scope has exited.")
			printClosure := `
			func main() {
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

			Result:`
			fmt.Println(printClosure)
			counter := closureFunc()
			fmt.Println(counter())
			fmt.Println(counter())
			fmt.Println(counter())
			fmt.Println(counter())
		case 2:
			//recursion
			fmt.Println("---\nRecursion")
			fmt.Println("Recursion in Go is when a function calls itself to solve a problem by breaking it down into smaller sub-problems, usually with a base case to stop the calls.")
			printRecursion := `
			// Recursion function of n + (n-1)

			func main() {
				fmtPrintln(recursionFunc(10))
			}

			func recursionFunc(n int) int {
				if n < 1 {
					return n
				}
				return n + recursionFunc(n-1)
			}

			Result:`
			fmt.Println(printRecursion)
			fmt.Println(recursionFunc(10))
		case 3:
			//Pointers
			fmt.Println("---\nPointers")
			fmt.Println("Pointers in Go are variables that store the memory address of another variable, allowing functions to modify values directly.")
			printPointer := `
			func main() {
	
			var ptr *int
			var a int = 10
			ptr = &a //referecing a pointer
				
			fmt.Println(a)
			fmt.Println(ptr)
			fmt.Println(&a)	  //same as ptr
			fmt.Println(*ptr) //dereference the pointer
			//a pointer that is not popointing to any value is a nil pointer (when no value is declared)
			if ptr == nil {
				fmt.Println("Pointer is nil")
			}

			modifyValue(ptr)
			fmt.Println(a)
			}

			func modifyValue(ptr *int) {	//when we pass a variable to a function, a copy of that variable is passed to the func.
				*ptr++			//but when we are using pointers, we are accessing the memory address of the var
							//and the value is modified. Thus no copy. :)
			}`
			fmt.Println(printPointer)
		case 4:
			fmt.Println("---\nStrings and Runes")

		default:
			fmt.Println("Exit")
			break loop
		}
		//Check if user wants to run the program agian or exit.
		var input string
		fmt.Print("Do you want to continue? (y/n): ")
		fmt.Scanln(&input)

		if input == "y" {
			fmt.Println("Continuing...")
		} else if input == "n" {
			fmt.Println("Exiting...")
			break // exit the loop
		} else {
			fmt.Println("Invalid input, please enter 'y' or 'n'")
		}

	}
}

func closureFunc() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Recursion function of n + (n-1)
func recursionFunc(n int) int {
	if n < 1 {
		return n
	}
	return n + recursionFunc(n-1)
}
