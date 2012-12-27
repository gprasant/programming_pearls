/*
x = generate continuous numbers from 0 to n
for i in [0, n-1)
	swap x[i] and x[rand(i, n-1)]
====================================
usage : go run gen_rand.go > random.txt 
*/
package main

import (
	"fmt"
	"math/rand"
)

var N int = 10000

func randlu(l, u int32) int32 {
	return l + rand.Int31n(u-l)
}

func get_slice() []int32 {
	x := make([]int32, N)
	for i := 0; i < N; i++ {
		x[i] = int32(i)
	}
	return x
}

func main() {
	x := get_slice()
	for i := 0; i < N; i++ {
		a := int32(i)
		b := randlu(a, int32(N))
		x[a], x[b] = x[b], x[a]
	}
	fmt.Println(x)
}
