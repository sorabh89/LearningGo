package main

import (
	"fmt"
)

func main() {
	//creating an array of size 4

	a := [4]int{1,2,3,4}

	fmt.Println(a)			//output : [1 2 3 4]
	fmt.Println(len(a))		//output : 4            length of the array

	//we can avoid providing the size by using ... as it'll calculate the size from the values provided

	b := [...]int{1,2,3,4,5}
	fmt.Println(b)			//output : [1 2 3 4 5]

	//Creating a 2D array

	var identityMatrix [3][3]int = [3][3]int{[3]int{0,0,1}, [3]int{0,1,0}, [3]int{1,0,0}}
	fmt.Println(identityMatrix)	//output : [[0 0 1] [0 1 0] [1 0 0]]

	//another variable to point to the same array using pointers
	c := &a
	fmt.Println(c)			//output : &[1 2 3 4]

	//changing the value in the pointing variable will also change it in the original array
	c[2] = 5
	fmt.Println(c)			//output : &[1 2 5 4]
	fmt.Println(a)			//output : [1 2 5 4]

	//Slices : slices are a chuck of an array and hence are always backed by an array.
	//creating a slice, will be backed by an array containing the values assigned
	//Syntax is similar to an array except the size or ...

	d := []int{1,2,3,4}
	fmt.Println(d)			//output : [1 2 3 4]

	//Length and capacity functions to print the length and capacity of an array/Slice.
	//In case of Slice since it is a slice of an array the length and capacity can be different, whereas in array It'll be same

	fmt.Printf("Length : %v, Capacity : %v \n", len(d), cap(d))		//output : Length : 4, Capacity : 4


	//Unlike array we don't need '&' pointers in slices, Slices are backed by an array and hence another pointer variable will also point to the same array
	e := d
	fmt.Println(e)			//output : [1 2 3 4]

	//changing the value in a slice is actually changing the main array and hence all the slices will be impacted
	e[2] = 5
	fmt.Println(e)			//output : [1 2 5 4]
	fmt.Println(d)			//output : [1 2 5 4]

	f := []int{0,1,2,3,4,5,6,7,8,9}
	g := f[:]
	h := f[3:]
	i := f[:6]
	j := f[3:6]

	fmt.Println(f)			//output : [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(g)			//output : [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(h)			//output : [3 4 5 6 7 8 9]
	fmt.Println(i)			//output : [0 1 2 3 4 5]
	fmt.Println(j)			//output : [3 4 5]

	//changing value of any slice will reflect for all the slices backed by that array
	h[2] = 42

	fmt.Println(f)			//output : [0 1 2 3 4 42 6 7 8 9]
	fmt.Println(g)			//output : [0 1 2 3 4 42 6 7 8 9]
	fmt.Println(h)			//output : [3 4 42 6 7 8 9]
	fmt.Println(i)			//output : [0 1 2 3 4 42]
	fmt.Println(j)			//output : [3 4 42]

	k := []int{}
	fmt.Printf("%v, Length : %v, Capacity : %v\n", k, len(k), cap(k))		//output : [], Length : 0, Capacity : 0

	//Slices provide the functionality to add or remove values and hence the backed arrays can change

	k = append(k, 1)
	fmt.Printf("%v, Length : %v, Capacity : %v\n", k, len(k), cap(k))		//output : [1], Length : 1, Capacity : 1


	//As we append g the slice will create another array of higher capacity (usually double the size)
	//copies the values of the underlying array and then append the value. Hence the change here in h will not reflect for g

	g = append(g, 10)
	h[1] = 41
	fmt.Println(f)			//output : [0 1 2 3 41 42 6 7 8 9]
	fmt.Println(g)			//output : [0 1 2 3 4 42 6 7 8 9 10]		//No 41 here bcoz its backed by a different array
	fmt.Println(h)			//output : [3 41 42 6 7 8 9]
	fmt.Println(i)			//output : [0 1 2 3 41 42]
	fmt.Println(j)			//output : [3 41 42]


	fmt.Printf("%v, Length : %v, Capacity : %v\n", g, len(g), cap(g))		//output : [0 1 2 3 4 42 6 7 8 9 10], Length : 11, Capacity : 20

	//append() can also be used to append more than one value

	k = append(k, 2, 3, 4)
	fmt.Printf("%v, Length : %v, Capacity : %v\n", k, len(k), cap(k))		//output : [1 2 3 4], Length : 4, Capacity : 4


	// use 3 dots (...) with the slice to append one slice to another, won't be able to append another array

	l := []int{5,6,7}
	k = append(k, l...)
	fmt.Printf("%v, Length : %v, Capacity : %v\n", k, len(k), cap(k))		//output : [1 2 3 4 5 6 7], Length : 7, Capacity : 8


	//Slices can be used for stacks and queues
	m := []int{1,2,3,4,5}

	//append() will add a value in the end
	//use slice operations to have values excluding the first or last element
	//However dropping a value from between is tricky

	n := append(m[:2], m[3:]...)		//The underlying array for n is also m, Hence this assignment is actually going to change m as well
	fmt.Println(n)				//output : [1 2 4 5]
	fmt.Println(m)				//output : [1 2 4 5 5]


	o := []int{1,2,3,4,5}
	p := append([]int{}, o[:2]...)		//Here we created p with a new array and hence the above change will be done on this new array
	p = append(p, o[3:]...)
	fmt.Println(p)				//output : [1 2 4 5]
	fmt.Println(o)				//output : [1 2 3 4 5]


	//we can also use make method to create a slice with higher capacity then the values it contains

	q := make([]int, 3)
	fmt.Printf("%v, Length : %v, Capacity : %v\n", q, len(q), cap(q))		//output : [0 0 0], Length : 3, Capacity : 3

	r := make([]int, 3, 100)
	fmt.Printf("%v, Length : %v, Capacity : %v\n", r, len(r), cap(r))		//output : [0 0 0], Length : 3, Capacity : 100

}
