package main

import (
	"fmt"
)

func main() {

	//creating an array in Go

	romanNums := map[string]int {
		"I": 1,
		"II": 2,
		"III": 3,
		"IV": 4,
		"V": 5,
	}
	fmt.Println(romanNums)						//output : map[I:1 II:2 III:3 IV:4 V:5]			values can be returned in any order

	//Slices cannot be used as a key in a map

	//sliceKeyMap := map[[]int]string{}				//Error : invalid map key type []int
	//fmt.Println(sliceKeyMap)

	//Unlike Slices, Arrays can be used as a key in maps
	arraysKeyMap := map[[2]int]string{}
	fmt.Println(arraysKeyMap)					//output : map[]

	//similar to arrays make() can also be used to create maps

	romanNums2 := make(map[string]int)
	romanNums2 = map[string]int {
		"I": 1,
		"II": 2,
		"III": 3,
		"IV": 4,
		"V": 5,
	}
	fmt.Println(romanNums2)						//output : map[I:1 II:2 III:3 IV:4 V:5]

	//getting a value from the map using the key

	fmt.Println(romanNums2["IV"])					//output : 4

	//Adding an element to the map

	romanNums2["VI"] = 6
	fmt.Println(romanNums2)						//output : map[I:1 II:2 III:3 IV:4 V:5 VI:6]

	//Deleting an element from the map

	delete(romanNums2, "IV")
	fmt.Println(romanNums2)						//output : map[I:1 II:2 III:3 V:5 VI:6]

	//When we try to get a value that doesn't exist in the map, it doesn't throw an error but returns the default value

	fmt.Println(romanNums2["X"])					//output : 0

	//for such situations when we need to find the key exists or not we have the ok statement

	ten, ok := romanNums2["X"]
	fmt.Println(ten, ok)						//output : 0 false

	romanNums2["X"] = 10
	ten, ok = romanNums2["X"]
	fmt.Println(ten, ok)						//output : 10 true


	//to find the number of elements in the map we can use the len() function

	fmt.Println(len(romanNums2))					//output : 6

	//Maps are passed by reference so if change is made using one variable, all variables pointing to the map will reflect it

	romans := romanNums2
	delete(romans, "III")

	fmt.Println(romanNums2)						//output : map[I:1 II:2 V:5 VI:6 X:10]
}
