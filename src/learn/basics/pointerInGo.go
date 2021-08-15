package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main() {
	a := 10
	b := a				//here b is another variable that has the same value as a but its a copy and not the same value
	fmt.Println(a, b)		//output :10 10
	a = 20
	fmt.Println(a, b)		//output :20 10		changing the value of a has no impact on b as both are 2 dfferent values


	var c *int = &a			//Creating a variable c of type 'pointer to an int' and using the address operator '&' to point it to the address of a

	//dereferencing operator (*)
	//to access the value stored at a pointer we can use the dereferencing variable as shown below

	fmt.Println(a, c, *c)		//output : 20 0xc000018050 20

	//c = 20			//Error: cannot use 20 (type int) as type *int in assignment

	//we'll need to use the dereferencing operator for opeartions like assigning a value or accessing the value
	*c = 30

	fmt.Println(a, c, *c)		//output : 30 0xc000018050 30

	arr := [3]int{1,2,3}
	d := &arr[0]
	e := &arr[1]

	//we can use '%p' to print the address of pointer variables
	fmt.Printf("%v %p %p\n", arr, d, e)	//output : [1 2 3] 0xc000014018 0xc000014020

	/*
	Note : Go has unsafe package for pointer related functions.
		However as the name hints the pointer functions should be used with utmost understanding as they deal with the memory and can be confusing
	*/

	//A pointer to an uninitialized variable has <nil> in it and hence can throw a panic in the program
	var ms *myStruct
	fmt.Println(ms)			//output : <nil>

	ms = new(myStruct)		//new can only initialize to a default value
	fmt.Println(ms)			//output : &{0}		default value

	(*ms).foo = 101 		//dereferencing operator has a lower precedence that dot(.) operator hence required to put it in brackets to execute it first

	fmt.Println(ms)			//output : &{101}
	fmt.Println((*ms).foo)		//output : 101

	//Instead of always using the dereferencing operator for pointer variables, Go compiler helps us by interpreting a pointer variable automatically without using the dereferencing variable
	fmt.Println(ms.foo)		//output : 101


	arr2 := [3]int{1,2,3}
	arr3 := arr2

	fmt.Println(arr2, arr3)		//output : [1 2 3] [1 2 3]
	arr2[1] = 22
	fmt.Println(arr2, arr3)		//output : [1 22 3] [1 2 3]		as in case of primitives an array also behaves in the same way arr3 has a copy of a arr2 and hence a change in arr2 has no impact on arr3

	/*
	However in case of slice its not the same,
	A slice is not an array but it points to an array and hence when we assign a slice to another variable we are still pointing to the same underlying array
	And so the changes using one variable will impact the other
	*/

	arr4 := []int{1,2,3}
	arr5 := arr4

	fmt.Println(arr4, arr5)		//output : [1 2 3] [1 2 3]
	arr4[1] = 22
	fmt.Println(arr4, arr5) 	//output : [1 22 3] [1 22 3]

	//The same is true for maps as well, as they are also pointing to the values

	m := map[string]string {"key1": "value1", "key2": "value2"}
	n := m
	fmt.Println(m, n)		//output : map[key1:value1 key2:value2] map[key1:value1 key2:value2]
	m["key1"] = "newValue"
	fmt.Println(m, n)		//output : map[key1:newValue key2:value2] map[key1:newValue key2:value2]
}
