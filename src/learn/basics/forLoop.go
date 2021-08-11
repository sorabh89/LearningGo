package main

import (
	"fmt"
)

func main() {

	// A simple for loop in Go

	for i := 0; i < 5; i++ {					//output : 0 1 2 3 4
		fmt.Print(i, " ")
	}

	fmt.Println("")

	//Unlike other languages Go handles more than one variable in a different way as shown below

	for i, j := 0, 5 ; i < 5 && j< 8; i,j = i+1, j+1 {			//output : 0,5 1,6 2,7
		fmt.Print(i, ",", j, " ")
	}

	fmt.Println("")

	// To initialize a variable outside the for loop, The syntax is as slown below

	i := 0
	for ; i < 5; i++ {					//output : 0 1 2 3 4
		fmt.Print(i, " ")
	}

	fmt.Println("")

	// Go doesn't have Do and while loops, Instead the for loop alone can do the work

	for i<10 {						//output : 5 6 7 8 9		//Also this is same as for ;i<10;
		fmt.Print(i, " ")
		i++
	}

	fmt.Println("")

	//We can also use for loop without any condition

	for {							//output : 10 11 12 13 14
		if i >= 15 {
			break
		}
		fmt.Print(i, " ")
		i++
	}

	fmt.Println("")

	//Label in loops

	for k:=1; k<=3; k++ {					//output : 1 2 3 2 4 3
		for l:=1; l<=3; l++ {
			if(k*l > 5) {				// without this break condition the output would be  1 2 3 2 4 6 3 6 9
				break				// But here the break is breaking from the inner for loop and not from the outer loop
			}					// To break from outer loop Go supports labels
			fmt.Print(k*l, " ")
		}
	}

	fmt.Println("")


	Breakhere:
	for k:=1; k<=3; k++ {					//output : 1 2 3 2 4
		for l:=1; l<=3; l++ {
			if(k*l > 5) {
				break Breakhere			// Here the break is breaking outside both the loops looking for the label 'Breakhere'
			}
			fmt.Print(k*l, " ")
		}
	}

	fmt.Println("")

	//Looping over collections :  To loop over collections we can use range that give the key and value. We can use it for arrays, slices, maps, string, etc

	slc := []int {1,2,3}

	for k,v := range slc {					//output : 0 1
		fmt.Println(k, v)				//	   1 2
	}							// 	   2 3
}
