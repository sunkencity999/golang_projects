package main

import "fmt"

func main() {
	//loop with if statements
	for i := 0; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Println(i, "\t", "Odd number")
		}
		if i%2 == 0 {
			fmt.Println(i, "\t", "Even number")
		}
	}
	//simple loop
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
	}

	//nested loop
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}

	//conditional loop with break. Do this action if the following action remains true and then break the loop
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}

	//conditional loop with continue and break. It will run the chunk of code before
	//continue until it gives a true result, after which it will continue to the next statement.main()
	l := 0
	for {
		l++
		if l%2 == 0 {
			continue //move on and execute next statement, otherwise return to top
		}
		fmt.Println(l)
		if l >= 20 {
			break //stop running the loop at this point.
		}
	}
}
