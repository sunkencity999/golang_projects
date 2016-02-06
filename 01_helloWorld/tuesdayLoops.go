package main

import "fmt"

func main() {
	//simple loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//loop with if statement
	for i := 0; i < 20; i++ {
		if i <= 5 {
			fmt.Println(i, "\t", "The value is less than 6.")
		}
		if i >= 6 {
			fmt.Println(i, "\t", "The value is 6 or greater.")
		}

	}

	//conditional loop with continue statement
	i := 0
	for {
		fmt.Println(i)
		i++
		if i < 10 {
			continue
		}
		break
	}

	// conditional loop with break statement
	j := 0
	for {
		fmt.Println(j)
		j++
		if j >= 10 {
			break
		}
	}

	//nested loop
	for k := 0; k <= 5; k++ {
		for l := 0; l <= 5; l++ {
			fmt.Println(k, "\t", l)
		}
	}
	//switch statement
	switch "Jimbo" {
	case "Tim":
		fmt.Println("Hello Tim")
	case "Christopher":
		fmt.Println("Hello Christopher!")
	default:
		fmt.Println("You are not the Christopher!")

	}

	//switch statement with fallthrough
	switch "Christopher" {
	case "Al":
		fmt.Println("Hey Juice")
	case "Heidi":
		fmt.Println("Hey Heid")
	case "Christopher":
		fmt.Println("Hey Christopher")
		fallthrough
	case "Family":
		fmt.Println("Hey Bradford Family!")
	}

	//switch statement with variable, will only print the case that evaluates to true --boolean
	mySwitchValue := "Chris"
	switch {
	case len(mySwitchValue) == 5:
		fmt.Println("This statement evaluated as true")
	case mySwitchValue == "Tim":
		fmt.Println(" This is wrong so it better not fucking print")
	}

}
