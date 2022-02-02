package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64 = 15046
	var result = math.Sqrt(num)
	var res = math.Round(result)
	var ceil = math.Ceil(result)
	var floor = math.Floor(result)
	fmt.Println(result)
	fmt.Printf("%.2f \n", result)
	fmt.Println(res)
	fmt.Println(floor)
	fmt.Println(ceil)
}
