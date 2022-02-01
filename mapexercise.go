package main

import "fmt"

func main() {
	s := "cabbage"
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		_, err := m[s[i]]
		if err == true {
			m[s[i]] = m[s[i]] + 1
		} else {
			m[s[i]] = 1
		}
	}
	var sum int
	for i, _ := range m {
		if i == 'a' || i == 'b' {
			sum = sum + 1*m[i]
		}
		if i == 'c' {
			sum = sum + 2*m[i]
		}
		if i == 'g' || i == 'e' {
			sum = sum + 3*m[i]
		}
	}
	/*for i, v := range m {
		fmt.Println(string(i), " value = ", v)
	}*/
	fmt.Println(sum)
}
