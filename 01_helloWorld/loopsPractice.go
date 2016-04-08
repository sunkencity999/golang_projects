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
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//loop with if statement
	for h := 0; h <= 5; h++ {
		if h == 5 {
			fmt.Println("You have found the H!")
		} else {
			fmt.Println("This number is not the H :-( )")
		}
	}

	//continue and break
	k := 0
	for {
		fmt.Println(k)
		k++
		if k < 5 {
			continue
		}
		break
	}

	//switch statement and taking user input
	var usrInput string
	fmt.Print("What is your name, User?: ")
	fmt.Scan(&usrInput)

	rslt := usrInput

	switch rslt {
	case "Christopher":
		fmt.Println("Hello, programmer!")
	default:
		fmt.Println("You are not the Programmer. Begone!")

	}

	//creating a slice using make
	ex := make([]int, 3)
	ex[0] = 1
	ex[1] = 2
	ex[2] = 3
	fmt.Println(ex)

	//creating a slice with shorthand
	ext := make([]int, 5)
	ext[0] = 1
	ext[1] = 2
	fmt.Println(ext)
	//deleting from a slice

	//creating a map (key/value pair)
	top := map[string]string{}
	top["Chris"] = "Programmer"
	top["Restayall"] = "Ain't shit"
	fmt.Println(top)

	var myMap = make(map[string]string)
	myMap["One"] = "Singular Sensation"
	myMap["Two"] = "hands, feet, eyes, lungs"
	fmt.Print(myMap)

	//using range to iterate over a map
	for key, value := range myMap {
		fmt.Println(key, " - ", value)
	}
	//using the range of an array
	p := check(2, 3, 4, 5)
	fmt.Println(p)
}
