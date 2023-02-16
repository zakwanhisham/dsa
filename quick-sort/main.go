package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)
	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}

func main() {
	slice := generateSlice(20)
	fmt.Println("\n---- Unsorted ---- \n\n", slice)
	quickSort(slice)
	fmt.Println("\n---- Sorted ---- \n\n", slice)
}
