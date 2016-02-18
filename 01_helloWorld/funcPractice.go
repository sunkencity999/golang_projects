package main

import "fmt"

func main() {
	greeting("")
	s := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 0
	})
	fmt.Println(s)
	fmt.Println(factorial(6))
}
func greeting(name string) {
	fmt.Print("What is your name? ")
	fmt.Scan(&name)
	fmt.Println("Welcome", name+"!")

	l := average(43, 15, 36, 87)
	fmt.Println(l)
	//var n float64
	//fmt.Println("Please input the numbers you would like the average of: ")
	//fmt.Scanln(&n...)
	//m := average(n)
	//fmt.Println(m)

}
func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func filter(numbers []int, callbackFunc func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callbackFunc(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

//example of recursion--a function calling upon itself. A factorial is a number that multiplies itself with all the numbers below it successively
func factorial(x int) int {
	if x == 0 {
		return 1 //if you return anything, the function stops there. The code
		//below will only run if x does not equal 0
	}
	return x * factorial(x-1) //this will take x and multiply at by x-1, continuing
	//to loop through itself until x == 0, at which point 1 is returned and it stops
}
