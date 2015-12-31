package main

import "fmt"

func main() {
	//fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	//prints 42 in decimal, than binary, thank base16/hexadecimal
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
		//for the beginning ' i := 0' is where we initialize a 'counter' variable named
		// 'i' that starts at/has the value of zero. Next we check a condition 'i, which has been initialized and given the
		// value of 0, is less than the value of 200' post (the section after the ; )
		// we instruct the program to do something after the loop code is run. in this
		// instance we say to increment to the next number.
	}
}
