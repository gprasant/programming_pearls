package main

import "fmt"

func main() {
	x := []int{3, 9, 5, 2, 1}
	quickSort(x, 0, 4)
	fmt.Println("After sort", x)
}

func swap(slice []int, idx1 int, idx2 int) {
	t := slice[idx1]
	slice[idx1] = slice[idx2]
	slice[idx2] = t
}

func quickSort(slice []int, l int, u int) {
	if l >= u {
		return
	}
	m := l
	for i := l + 1; i <= u; i++ {
		if slice[i] < slice[m] {
			m++
			swap(slice, m, i)
		}
	}
	swap(slice, l, m)

	quickSort(slice, l, m-1)
	quickSort(slice, m+1, u)
}
