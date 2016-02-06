package main

import "fmt"

func main() {
	p := 0
	for p < 10 {
		if p <= 4 {
			fmt.Println(p, "\t", "is equal to or less than four.")
		} else if p >= 6 {
			fmt.Println(p, "\t", "is equal to or greater than six.")
		} else {
			fmt.Println(p, "\t", "is equal to five.")
		}
		p++
	}

}
