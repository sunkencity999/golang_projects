package main

import "fmt"

func main() {
	var input string
	fmt.Print("Please enter your name:")
	fmt.Scan(&input)
	fmt.Println("Hello", input+"!")

}
