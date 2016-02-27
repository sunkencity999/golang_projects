package main

import "fmt"

func check(n ...int) []int {
	var tote []int
	for _, k := range n {
		tote = append(tote, k)
	}
	return tote
}

func main() {
	//simple loop
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//loop with if statement
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println(i)
		} else {
			fmt.Println("This ain't 5.")
		}
	}

	//continue and break
	k := 0
	for {
		fmt.Println(k)
		k++
		if k >= 10 {
			continue
		}
		break
	}

	//switch statement
	var usrInput string
	fmt.Print("Please input your name: ")
	fmt.Scan(&usrInput)
	taken := usrInput

	switch taken {
	case "Christopher":
		fmt.Println("Welcome, Programmer")
	default:
		fmt.Println("You are not the programmer. Begone, Peasant!")
	}

	//taking user input

	//creating a slice using make
	list := make([]int, 3, 5)
	list[0] = 7
	list[1] = 8
	list[2] = 9
	for run := range list {
		fmt.Println(run)
	}
	//creating a slice with shorthand
	listTwo := []int{1, 2, 3, 4}
	listThree := []int{5, 6, 7, 8}

	listTwo = append(listTwo, listThree...)
	fmt.Println(listTwo)

	//using the range of an array
	p := check(2, 3, 4, 5)
	fmt.Println(p)
}
