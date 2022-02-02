package main

import (
	"errors"
	"fmt"
)

func pal(n int) bool {
	var rem, num, temp int
	temp = n
	for {
		rem = n % 10
		num = num*10 + rem
		n = n / 10
		if n == 0 {
			break
		}
	}
	if temp == num {
		return true
	} else {
		return false
	}
}

func fac(n int) (num int, er error) {
	fact := 1
	if n <= 0 {
		er = errors.New("Enter valid number")
		return num, er
	} else {
		for i := 1; i <= n; i++ {
			fact = fact * i
		}
		return fact, er
	}
}

func main() {
	var n int
	fmt.Println("Enter your number")
	fmt.Scanf("%d", &n)
	fmt.Println(pal(n))
	f, e := fac(0)
	if e == nil {
		fmt.Println(f)
	} else {
		fmt.Println(e)
	}
}
