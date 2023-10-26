/*This exercise determines the diagonal difference in a square matrix array. [][]int32

Example:

1,2,3
8,2,6
7,9,2

1 + 2 + 2 = 5
7 + 2 + 3 = 12

absolute difference = 5 - 12 = 7 (no negatives)

*/

package main

import "fmt"

func main() {

	//define our variables
	var primary int32 = 0
	var secondary int32 = 0

	var row1 = []int32{1, 4, 2}
	var row2 = []int32{1, 3, 2}
	var row3 = []int32{4, 3, 2}

	var arr = [][]int32{row1, row2, row3}

	//loop over the array
	//primary increments by e[i]
	//[1,2,3]
	//[2,3,4]
	//[5,6,7]

	//loop 1 = e[0] = 1 (index 0)
	//loop 2 = e[1] = 3 (index 1)
	//loop 3 = e[2] = 7 (index 2)

	//secondary basically does this in reverse
	//same array as above example
	//loop 1 = e[(length of e(3) - 0)-1] = index 2
	//loop 2 = e[(length of e(3) - 1)-1] = index 1
	//loop 3 = e[(length of e(3) - 2)-1] = index 0
	for i, e := range arr {
		primary += e[i]
		secondary += e[(len(e)-i)-1]
	}

	fmt.Println(primary)
	fmt.Println(secondary)
	fmt.Println(getAbsolute(primary - secondary))
}

func getAbsolute(num int32) int32 {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
