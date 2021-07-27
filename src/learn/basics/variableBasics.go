package main

import (
	"fmt"
	"strconv"
)

var i float32 = 10		            	//lowercase variables are scoped to the package level and can't be accessed outside the package
//var I float32 = 10				        //Need to have variable with first letter capital to access it outside the package

var (                               //Declaring a set of variables
  s1 = 10
  s2 = "this is a string"
)

func main() {
	fmt.Printf("%v, %T \n", i, i)		  //output : 10, float32

	var i int = 21			              //shadowing variable i declared at package level
	//i := 22
	fmt.Printf("%v, %T \n", i, i)		  //output : 21, int

	s3 := 21				                  // auto identifies a variable and its type
	fmt.Println(s3)				            //output : 21

	var j float32 = 50.5

	var k = int(j)
	fmt.Printf("%v, %T \n", k, k)		  //output : 50, int

	var l string = string(42)
	fmt.Printf("%v, %T \n", l, l)		  //output : *, string		unicode value

	var m string = strconv.Itoa(42)
	fmt.Printf("%v, %T \n", m, m)		  //output : 42, string

  	var n int32 = 4
	//fmt.Println(i+n)			          //throws error: invalid operation: i + n (mismatched types int and int32)
	fmt.Println(i + int(n))			      //output : 25      need to have same types

	fmt.Printf("%v, %T \n", s2[0], s2[0])		        //output : 116, uint8      string can be used as an array but the value is treated as a byte so explicit conversion in required
	fmt.Printf("%v, %T \n", string(s2[0]), s2[0])   //output : s, uint8

	//s2[0] = "a"				              //throws error: cannot assign to s2[0] (strings are immutable)

	b := []byte(s2)
	fmt.Printf("%v, %T \n", b, b)		  //[116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8      converts string to byte array

	var r rune = 'a'			            //runes are int32 types, utf32 characters
	fmt.Printf("%v, %T \n", r, r)		  //output : 97, int32
}
