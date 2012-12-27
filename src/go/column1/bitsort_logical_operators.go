/*
test - map[i/32] & (1 << (i & 1F))
set  - map[i/32] |= (1 << (i & 1F))
clr  - map[i/32] &^= (1 << (i & 1F)) [In Go, &^ is the bitwise clear operator]
*/
package main

import "fmt"

var N int = 8 // for this program to work, N should be a power of 2
var bitmap = make([]int32, N/32+1)

var WORD_LENGTH int32 = 32
var MASK int32 = WORD_LENGTH - 1

func main() {
	var num int

	for i := 0; i < N; i++ {
		fmt.Scanf("%d\n", &num)
		set(int32(num))
	}

	for i := 0; i < N; i++ {
		if test(int32(i)) {
			fmt.Println(i)
		}
	}
}

func test(i int32) bool {
	return (bitmap[i/WORD_LENGTH]&(1<<uint((i&MASK))) != 0)
}

func set(i int32) {
	bitmap[i/WORD_LENGTH] |= (1 << uint((i & MASK)))
}

func clr(i int32) {
	bitmap[i/WORD_LENGTH] &^= (1 << uint((i & MASK)))
}
