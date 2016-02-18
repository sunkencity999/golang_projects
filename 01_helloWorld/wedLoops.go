package main

import "fmt"

func main() {
	//basic for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//loop with continue statement
	i := 0
	for {
		fmt.Println(i)
		i++
		if i < 10 {
			continue
		}
		break
	}

	//loop with break statement
	j := 0
	for {
		fmt.Println(j)
		j++
		if j >= 10 {
			break
		}
	}

	//loop with if statements
	for k := 0; k < 10; k++ {
		if k < 5 {
			fmt.Println(k, "\t", "This value is less than five.")
		}
		if k > 5 {
			fmt.Println(k, "\t", "This value is greater than five.")
		}
	}

	//nested loop printing multi-layered results
	for m := 0; m < 5; m++ {
		for n := 0; n < 3; n++ {
			fmt.Println(m, "\t", n)
		}
	}
	//case-switch statement
	switch "Christopher" {
	case "Tim":
		fmt.Println("This is not Christopher. You are terrible")
	case "Christopher":
		fmt.Println("You are the true Christopher. Congratulations!")
	}

	//if-else if-else statement loop
	p := 0
	for p < 10 {
		if p <= 4 {
			fmt.Println(p, "\t", "is equal to or less than four.")
		}
		if p >= 6 {
			fmt.Println(p, "\t", "is equal to or greater than six.")
		}
		if p == 5 {
			fmt.Println(p, "\t", "is equal to five.")
		}
		p++
	}

}
