package main

import (
	"fmt"

//	"strings"
)

func main() {

	var k, m int
	var a, c []int
	fmt.Println("Enter k : $ ")
	fmt.Scanf("%d\n", &k)
	fmt.Println("Enter m : $")
	fmt.Scanf("%d\n", &m)
	// input
	c = []int{1, 2, 3, 4} // for k = 3
	a = []int{1, 2, 3}
	/* FAILED USER INPUT : Need to get a and c from User
	fmt.Println("Enter c1, c2 .. ck (each Separated by spaces) $")

	var s string
	_, err := fmt.Scanf("%s\n", &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("s = ", s)
	c = strings.Split(s, " ")

	fmt.Println("Enter a1, a2 .. ak (each Separated by spaces)\n $")

	fmt.Scanf("%s\n", &s)
	a = strings.Split(s, " ") */

	for n := k; n < m; n++ {
		a = append(a, c[k])
		for x := 0; x < k; x++ {
			a[n] += c[x] * a[n-x-1]
		}
	}

	fmt.Println(a)
}
