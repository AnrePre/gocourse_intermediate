package main


import "fmt"


func main() {
	sequence:= adder()
	fmt.Println(sequence())

}

func adder () func() int {
	i := 0
	fmt.Println("Previous value of i:", i)
	return func() int{
		i++
		fmt.Println("added 1 to i")
		return i
	}
	

}
