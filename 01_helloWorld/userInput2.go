package main

import "fmt"

func main() {
	var inputSmall int //create a variable to hold an integer
	var inputLarge int //create second vaiable to hold an integer

	fmt.Print("Please enter a small number:")
	fmt.Scan(&inputSmall) //use 'fmt.Scan' to capture stdin (user input) and store in the variable using '&' operator
	fmt.Print("Please enter a larger number:")
	fmt.Scan(&inputLarge)                      //use 'fmt.Scan' to capture stdin (user input) and store in the variable using '&' operator
	result := inputLarge % inputSmall          // store result of operation taking the remainder of inputLarge divided by inputSmall in 'result' variable
	fmt.Println("Here is your result", result) //print result

}
