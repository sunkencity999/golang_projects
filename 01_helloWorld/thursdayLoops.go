package main

import "fmt"

func main() {
	/*	for i := 0; i < 101; i++ {
			if i%3 == 0 {
				fmt.Println("Fizz")
			} else if i%5 == 0 {
				fmt.Println("Buzz")
			} else if i%3 == 0 && i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println(i)
			}
		}
	*/
	counter := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println(counter)
}
