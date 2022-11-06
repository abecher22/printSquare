package main

import "fmt"

func main() {
	printSquare("*", 4)
	printSquare("=", 10)
	printSquare("Yada ", 2)
	printSquare("repeat", 20)
}

// printSquare prints a square of the given size using the given string
func printSquare(c string, size int) {
	// print the first line
	printFullLine(c, size)

	// print a line that starts and ends with c
	for i := 0; i < size-2; i++ {
		printMiddleLine(c, size)
	}

	// print the last line
	printFullLine(c, size)
}

// printFullLine prints a line that looks like "****"
func printFullLine(c string, size int) {
	// loop up to "size" time and print the string in "c" each time
	for i := 0; i < size; i++ {
		fmt.Printf("%s", c)
	}
	// print a new line
	fmt.Printf("\n")
}

// printMiddleLine prints a line that looks like "*   *"
func printMiddleLine(c string, size int) {
	// print the first "c" string in the line
	fmt.Printf("%s", c)
	// print spaces size-2 times
	for i := 0; i < size-2; i++ {
		fmt.Print(" ")
	}
	// print the last "c" string in the line
	fmt.Printf("%s\n", c)
}
