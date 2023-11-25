package main

import "fmt"

/*

Loop through 1000 numbers printing each number.

If a number is divisible by 3, replace with "Fizz"
If a number is divisible by 5, replace with "Buzz"
If a number is divisible by both 3 and 5, replace with "FizzBuzz"

*/

func main() {
	output := ""
	for i := 1; i <= 1000; i++ {

		if i%3 == 0 && i%5 == 0 {
			output += "FizzBuzz"
		} else if i%3 == 0 {
			output += "Fizz"
		} else if i%5 == 0 {
			output += "Buzz"
		} else if output == "" {
			output = fmt.Sprint(i)
		}
		fmt.Println(output)
		output = ""
	}
}
