package main

import "fmt"

func main() {
	unordered := []int{12, 11, -4, 2, 98, 13, 5, 6, 7}
	var low, high int
	low = 0
	high = len(unordered) - 1
	fmt.Println("Unordered array = ", unordered)
	quicksort(unordered, low, high)
	fmt.Println("Ordered array =", unordered)

}

func quicksort(ar []int, low int, high int) {
	if low < high {

		p := partition(ar, low, high)
		quicksort(ar, 0, p-1)
		quicksort(ar, p+1, high)
	}
}

func partition(ar []int, low int, high int) int {
	var pivot, temp int
	pivot = ar[high]
	i := low
	for j := low; j < high; j++ {
		if ar[j] < pivot {
			temp = ar[i]
			ar[i] = ar[j]
			ar[j] = temp
			i++
		}
	}
	ar[high] = ar[i]
	ar[i] = pivot

	return i
}
