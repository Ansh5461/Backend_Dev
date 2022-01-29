package main

import "fmt"

type Pill int

const (
	Placebo       Pill = iota // iota =0,1,2,3,4,5
	Aspirin                   //1
	Ibuprofen                 //2
	Paracetamol               //3
	Acetaminophen             //4
)

/*
var(
	name string
	age int
)
*/

/* enter    go get golang.org/x/tools/cmd/stringer
then enter stringer -type=Pill
*/
func main() {
	fmt.Println(Paracetamol)
}
