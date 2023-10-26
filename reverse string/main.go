package main

import (
	"fmt"
	"time"
)

func main() {
	str := "Reverse This String!"
	fmt.Println(ReverseString(str))
}

func ReverseString(str string) string {
	start := time.Now()
	//declare our new string
	newStr := ""
	/* Loop through our string we pass in the function
	Get the reverse of the current i and add that to newstr */
	for i := range str {
		newStr += string(str[(len(str)-i)-1])
	}

	fmt.Println("Reverse String execution time -", time.Since(start))

	return newStr
}
