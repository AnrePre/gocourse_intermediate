package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func() {
		fmt.Println("Closing opened file.")
		file.Close()
	}()
	fmt.Println("File opened successfully.")

	// Read the contents of the file
	data := make([]byte, 1024) // Buffer to read data
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error reading data from file:", err)
		return
	}

	fmt.Println("File content:", string(data))

}
