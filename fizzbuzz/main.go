/*
Print "FizzBuzz" if a number is divisible by both 3 and 5. If the number is divisible only by 3,
print "Fizz." If the number is divisible only by 5, print "Buzz." For numbers that are not
divisible by either, print the number itself.
*/
package main

import "fmt"

var count int = 15

func fizzBuzz(x int) {
	for i := 1; i <= x; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	var num int
	fmt.Println("Please enter the value: ")
	fmt.Scanln(&num)
	fizzBuzz(num)
}
