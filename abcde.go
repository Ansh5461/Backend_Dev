package main

import "fmt"

/*func main() {
	s := "abcdef"
	var c string
	c = "true"
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)-1; j++ {
			if s[i] == s[j] {
				c = "false"
				break
			}
		}
	}
	fmt.Println(c)
}
*/
func main() {

	var s string
	fmt.Println("enter string you want to check as an isogram")
	fmt.Scanln(&s)
	CheckIsogram(s)

}

func CheckIsogram(s string) {
	var c string
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)-1; j++ {
			if s[i] == s[j] {
				c = "false"
			}
		}
	}
	if c == "false" {
		fmt.Println("the string is not an isogram")
	} else {
		fmt.Println("the string is an isogram")
	}

}
