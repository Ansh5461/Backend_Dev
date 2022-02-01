package main

import "fmt"

func maxelement(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n int
	var l, r, j, i int
	s := []int{3, 0, 2, 0, 4}
	n = len(s)
	var wc int

	for i = 0; i < n; i++ {
		l = s[i]
		for j = 0; j < i; j++ {
			l = maxelement(l, s[j])
		}
		r = s[i]
		for j = i + 1; j < n; j++ {
			r = maxelement(r, s[j])
		}
		wc = wc + (min(l, r) - s[i])
	}
	fmt.Println("Water contained = ", wc)

}
