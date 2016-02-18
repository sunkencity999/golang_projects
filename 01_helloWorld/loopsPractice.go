package main

import "fmt"

func main() {
	//simple loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//loop with if statement
	for l := 0; l < 10; l++ {
		if l < 5 {
			fmt.Println(l, "\t", "Is less than five.")
		}
		if l == 5 {
			fmt.Println(l, "\t", "Is equal to five.")
		}
		if l > 5 {
			fmt.Println(l, "\t", "is greater than five")
		}

	}
	//continue and break
	k := 0
	for {
		fmt.Println(k)
		k++
		if k < 10 {
			continue
		}
		break
	}

	//switch statement
	var input string
	fmt.Print("Please enter your name, programmer: ")
	fmt.Scan(&input)
	switch input {
	case "johnathon":
		fmt.Println("Hello jonathon! I guess you aight.")
	case "Christopher":
		fmt.Println("Great name! Well done, operator.")
	}
	if input != "Christopher" {
		fmt.Println("You are not Christopher. You suck. ")
	}

	//taking user input

}
