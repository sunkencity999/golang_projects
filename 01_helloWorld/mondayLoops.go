package main

import "fmt"

func main() {
	//loop with if statement
	for i := 0; /*initialize*/ i < 10; /*condition*/ i++ /*post action*/ {
		if i < 5 {
			fmt.Println(i, "\t", []byte("is less than five."))
		}
		if i == 5 {
			fmt.Println(i, "\t", "is equal to five.")
		}
		if i > 5 {
			fmt.Println(i, "\t", "is greater than five.")
		}
	}
	//simple loop
	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}
	//nested loop
	for i := 1; i < 11; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, "\t", j)
		}
	}
	j := 0 //conditional loop with continue statement
	for {
		fmt.Println(j)
		j++
		if j < 10 {
			continue
		}

		break
	}
	k := 0 //conditional loop using break
	for {
		fmt.Println(k)
		k++
		if k >= 10 {
			break
		}
	}

}
