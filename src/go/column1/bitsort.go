/*
initialize bitmap to zero
foreach number in inpuut
	read number and set the corresponding element in bitmap to 1
iterate over bitmap and print i if bitmap[i] is true
*/
package main

import "fmt"

var N int = 10

func main() {
	var num int
	bitmap := make([]bool, N)

	for i := 0; i < N; i++ {
		fmt.Printf("Debug: i : %d ; N : %d \n", i, N)
		fmt.Scanf("%d", &num)
		bitmap[num] = true
	}

	for i, v := range bitmap {
		if v {
			fmt.Println(i)
		}
	}
}
