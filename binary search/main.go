package main

import "fmt"

func main() {

	//initialize our slice of ints
	arr := []int{1, 3, 15, 20, 22, 33, 44, 81, 200, 210}

	//check our result and determine if the key was found
	key := BinarySearch(arr, 0, len(arr)-1, 20)
	if key != -1 {
		fmt.Println("Number found at index:", key)
	} else {
		fmt.Println("Number not found in that array.")
	}
}

func BinarySearch(arr []int, start int, end int, key int) int {

	//for loop
	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == key {
			return mid
		}
		if arr[mid] < key {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	//if key is not found, return with -1
	return -1
}
