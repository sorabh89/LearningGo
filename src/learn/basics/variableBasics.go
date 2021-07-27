package main

import (
	"fmt"
	"strconv"
)

var i float32 = 10		//lowercase variables are scoped to the package level and cant be accessed outside the package

func main() {
	fmt.Printf("%v, %T \n", i, i)		//output : 10, float32

	var i int = 21			//shadowing variable i declared at package level
	//i := 22
	fmt.Printf("%v, %T \n", i, i)		//output : 21, int

	var j float32 = 50.5

	var k = int(j)
	fmt.Printf("%v, %T \n", k, k)		//output : 50, int

	var l string = string(42)
	fmt.Printf("%v, %T \n", l, l)		//output : *, string		unicode value

	var m string = strconv.Itoa(42)
	fmt.Printf("%v, %T \n", m, m)		//output : 42, string
}
