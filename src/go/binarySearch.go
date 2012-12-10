package main

import (
	"fmt"
)
var t int
var x = []int{1, 2, 3, 4, 5}
func main() {
	
	fmt.Println("Enter val to find ")
	fmt.Scanf("%d", &t)
	fmt.Println(BinarySearch(t, x))
}

func BinarySearch(t int, x []int) int {
	l := 0
	u := len(x) - 1
	for ;l <= u; {
		m := (l + u)/2
		if(x[m] == t){
			return m
		} else if(x[m] < t){
			l = m + 1
		} else if(x[m] > t){
			u = m - 1
		}
	}
	return -1
}