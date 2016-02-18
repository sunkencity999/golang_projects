package main

import "fmt"

func main() {
	var userInput int
	fmt.Println("Please Enter the number you wish to discover the Largest prime factor for:")
	fmt.Scan(&userInput)
	var i int
	i += userInput
	c := i / 2
	if i%2 == 0 {
		fmt.Println(c)
		i := c / 2
		fmt.Println(i)
	}

}
