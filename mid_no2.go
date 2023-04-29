// 2.	Using the concept of Array and looping in GoLang, create an array checking program to compare the contents of 2 arrays and explain each step!

// answer
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// The var input string statement declares a string variable input.
	var input string
	// Create a new Scanner object that reads from the standard input (os.Stdin) and assigns it to the variable scanner.
	scanner := bufio.NewScanner(os.Stdin)
	// Declare two []string variables named arr1 and arr2
	var arr1, arr2 []string

	// Print "array 1:" to the console.
	fmt.Println("array 1:")
	// Read the input from the standard input using the scanner object.
	scanner.Scan()
	// Assign the text that was read by the scanner to the input variable.
	input = scanner.Text()
	// Split the input string into an array of strings using strings.Split() and assigns it to arr1.
	arr1 = strings.Split(input, " ")

	// goes the same for array 2
	fmt.Println("array 2:")
	scanner.Scan()
	input = scanner.Text()
	arr2 = strings.Split(input, " ")

	// Initialize a loop that will iterate over each index of the arr1 slice.
	for i := 0; i < len(arr1); i++ {
		// check if the value at the current index of arr1 is different from the value at the same index of arr2.
		if arr1[i] != arr2[i] {
			// If different, print "index ke i berbeda" to the console, where i is the current index.
			fmt.Println("index ke", i, "berbeda")
		}
	}
	// o.	After the loop has finished iterating, the program will exit.
}

