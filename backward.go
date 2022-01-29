package main

import "fmt"

func main() {
	s := "GEEKS FOR GEEKS"
	m := map[byte]int{}
	var str string
	for i := len(s) - 1; i >= 0; i-- {
		//start from back
		_, err := m[s[i]]
		//check if element already there
		if !err {
			//if yes then put its value as zero
			m[s[i]] = 0
		}
		if m[s[i]] == 0 {
			if s[i] == ' ' {
				continue
			}
			str = str + string(s[i])
			m[s[i]] = m[s[i]] + 1
		}
		if m[s[i]] > 0 {
			continue
		}
	}
	fmt.Println(str)
}
