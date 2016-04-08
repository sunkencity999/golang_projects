package main

import "fmt"

//Numlist will create the list of prime numbers up to the ideal.
func NumList(max int, ch chan<- int) {
	ch <- 2
	for i := 3; i <= max; i += 2 {
		ch <- i
	}
	ch <- -1 //channel signal that the end/limit has been reached
}

//Filter will filter out the non-prime results
func Filter(in <-chan int, out chan<- int, prime int) {
	for i := <-in; i != -1; i = <-in {
		if i%prime != 0 {
			out <- i
		}
	}
	out <- -1 //channel signal that he end has been reached
}

//CalcPrime will return the prime factors of the user input
func CalcPrime(num int) []int {
	k := []int{}
	ch := make(chan int)

	go NumList(num, ch)
	for prime := <-ch; (prime != -1) && (num > 1); prime = <-ch {
		for num%prime == 0 {
			num = num / prime
			k = append(k, prime)
		}

		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	return k
}

func main() {
	var userInput int
	fmt.Println("Please Enter the number you wish to discover the Largest prime factor for:")
	fmt.Scan(&userInput)

	fmt.Println(CalcPrime(userInput))

}
