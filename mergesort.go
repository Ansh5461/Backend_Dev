package main

import "fmt"

//mergesort in go
func main() {
	unsorted := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Array before sorting ", unsorted)
	sorted := mergesort(unsorted)
	fmt.Println("Array after sorting ", sorted)
}

func mergesort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := mergesort(arr[len(arr)/2:])
	second := mergesort(arr[:len(arr)/2])
	final := merge(first, second)

	return final
}

func merge(m []int, n []int) []int {
	var i, j int
	ar := []int{}
	for i < len(m) && j < len(n) {
		if m[i] < n[j] {
			ar = append(ar, m[i])
			i++
		} else {
			ar = append(ar, n[j])
			j++
		}
	}
	for ; i < len(m); i++ {
		ar = append(ar, m[i])
	}
	for ; j < len(n); j++ {
		ar = append(ar, n[j])
	}
	return ar
}
