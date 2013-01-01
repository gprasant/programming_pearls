/*
This is different from the original binary search in that it returns the first occurence of an element in a sorted array
*/
package main

import (
	"fmt"
)

func main() {
	var t int
	x := []int{1, 2, 2, 4, 5}
	fmt.Print("Enter t : $")
	fmt.Scanf("%d\n", &t)
	l := -1
	u := len(x)

	for l+1 != u {
		m := (l + u) / 2
		if x[m] < t {
			l = m
		} else {
			u = m
		}
	}

	p := u
	if p == len(x) || x[p] != t {
		p = -1
	}
	fmt.Println(p)
}
