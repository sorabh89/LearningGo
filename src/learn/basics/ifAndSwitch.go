package main

import (
	"fmt"
)

func main() {

	romanNums := map[string]int {
		"I": 1,
		"II": 2,
		"III": 3,
		"IV": 4,
		"V": 5,
	}

	//In Go, Unlike othr languages if condition supports initialization along with the condition

	if pop, ok :=romanNums["III"]; ok {
		fmt.Println(pop)						//output : 3
	}

	if i:=1; i<5 {
		fmt.Println(i)							//output : 1
	}

	//Switch statement : Similar to if switch also supports initialization as shown below, And we can compare it with multiple values in the same case unlike in other languages e.g. Java

	switch j:= 4 + 4; j {							//output : something else
		case 1, 3, 5 : fmt.Println("1, 3, 5")
		case 2, 4, 6 : fmt.Println("2, 4, 6")
	 	default : fmt.Println("something else")
	}


	//Go also supports switch case without tag, Also Go doesn't need explicit break statement to break, Go implicitly supports break

	k:= 10

	switch {								//output : Less than or equal to 10
		case k <= 10 :
				fmt.Println("Less than or equal to 10")
		case k <= 20 :
				fmt.Println("Less than or equal to 20")
		default :
				fmt.Println("Greater than 20")
	}

	//As Go implicitly supports break, In case if we want to fallthrough to the next case we can use 'fallthrough' command

	switch {								//output : Less than or equal to 10
		case k <= 10 : 							//	   Less than or equal to 20
				fmt.Println("Less than or equal to 10")
				fallthrough
		case k <= 20 :
				fmt.Println("Less than or equal to 20")
		default :
				fmt.Println("Greater than 20")
	}


	//Type switch : Go also supports switches on the basis of types as shown below

	var l interface{} = 'a'
	switch l.(type) {							//output : Its something else
		case int :
				fmt.Println("Its an int")
				break
				fmt.Println("This line will not be printed")
		case float64 :
				fmt.Println("Its an float64")
		case string :
				fmt.Println("Its an string")
		default :
				fmt.Println("Its something else")
	}


}
