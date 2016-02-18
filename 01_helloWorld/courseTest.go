package main

import "fmt"

// write a function which takes an integer, has two returns--argument divided by 2
// and bool value. This is my answer below was imperfect, as instead of "returning"
//the values I printed one of them, and ended up with an extra return. This is because I didn't
//fully understand how to reference multiple types and parameters in a single function.
func isitTrue(num int) bool {
	n := num / 2
	if n%2 == 0 {
		fmt.Println(n, "True")
		return true
	}
	if n%2 == 1 {
		fmt.Println(n, "False")
		return false
	}
	return false
}

//this is the correct function. Takes an integer as input, returns an interger and a bool.
// the return statment is separted by a comma to differentiate between the two returns; to
//return a bool we right a statement that evaluates to true or false.
func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

//did this one mostly right--just failed to declare the type the first time and tried to use []int instead of just int.
func varied(k ...int) int {
	var big int
	for _, v := range k {
		if v > big {
			big = v
		}
	}
	return big
}

func foo(num ...int) int {
	var sum int
	for _, m := range num {
		sum += m
	}
	return sum
}

func main() {
	fmt.Println(isitTrue(7))

	//the below func expression needs to be built within func main, as you cannot declare
	//variables and make statements outside of functions. It does the same things as the
	//function 'half', but does it using anonymous functions.
	halfer := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	//this initializes two variables with the value of the single function "half", necessary
	//because we need a value to hold each return from this function. Two returns, two values.
	h, even := half(5)
	fmt.Println(h, even)
	l, odd := halfer(6)
	fmt.Println(l, odd)
	s := varied(4, 18, 24, 36)
	fmt.Println(s)
	k := foo(1, 2, 3, 4)
	fmt.Println(k)
	aSlice := []int{1, 2, 3, 4}
	fmt.Println(foo(aSlice...))

}
