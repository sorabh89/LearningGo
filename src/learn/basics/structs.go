package main

import (
	"fmt"
	"reflect"
)


//creating a struct

type Employee struct {
	Id int							  //Id, Name are accessible from outside the package as they start with a capital letter
	Name string	`required max:"100"`			//whereas techs is not accessible outside the package
	techs []string
}

type SoftwareEngineer struct {
	emp Employee
	country string
}

func main() {

	//instantiating a struct

	emp1 := Employee {
		Id : 1,
		Name : "John",
		techs : []string {"Java", "Go"},		//remember to put a comma in the end else the program will throw a compile time error
	}

	fmt.Println(emp1)					//output : {1 John [Java Go]}

	//structs can also be instantiated with positions, However not prefered as if a struct is changed it can impact the existing instances

	emp2 := Employee {
		2,
		"David",
		[]string {"Go", "Python"},
	}

	fmt.Println(emp2)					//output : {2 David [Go Python]}


	//Anonymous structs: Incase we need a struct that is short lived or we dont need multiple instances of it, then we can go for anonymous structs

	worker1 := struct{
		id int
		name string}{
		id : 1,
		name : "Johnny",
		}
	fmt.Println(worker1)					//output : {1 Johnny}

	//Unlike Maps and Slices, Structs are pass by value

	worker2 := worker1
	worker2.name = "Mak"

	fmt.Println(worker1)					//output : {1 Johnny}
	fmt.Println(worker2)					//output : {1 Mak}

	//To pass by reference we can use & to point to the same value

	worker3 := &worker2
	worker3.name = "Silva"

	fmt.Println(worker2)					//output : {1 Silva}
	fmt.Println(worker3)					//output : &{1 Silva}


	//Go doesn't support inheritance but to construct complex datatypes we can use structs in a has-a relation

	se := SoftwareEngineer {
		emp : Employee{3,
		"Randal",
		[]string {"Html", "Python"},},
		country : "Singapore",

	}

	fmt.Println(se)						//output : {{3 Randal [Html Python]} Singapore}
	fmt.Println(se.emp.Name)				//output : Randal


	//Tags in Structs
	// `required max:"100"` mentioned above is a tag used to specify some information about the field, however this is useless for Go unless a library uses it for some validation

	t := reflect.TypeOf(emp1)
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)					//output : required max:"100"
}
