package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//defer
	//funcDefer()
	//httpCallExample()

	//panic
	//divideByZero()
	//deferAndPanic()

	/*
	Recover
	So as we have seen once a panic happens the execution of the program (entire stack trace) stops after all the defers in the queue are executed.
	But what if we don't want the entire program to stop, we may still want to continue with the remaining functionality.
	For such scenarios Go provides recover(), it help the program to continue execution after recovering from the panic happened.
	However, remember that the remaining statements in the function in which panic happened, will not be executed unlike catch() in Java,
	Hence, Make sure that the functions doesn't contain functionality not related or dependent on the reason for panic.
	*/
	panicAndRecover()
}



/*
Defer are calls that are executed in the end of a function before it returns any value
Defer are called in the LIFO order, So the last defer statement will be called first
output :
Start  1
Middle  3
End  5
Defer 3  6
Defer 2  4
Defer 1  2

Note : Though the defer statements are called in the end, the value of i is same as you read the code without defer.
This is because, when a defer is encountered it is queued to be called in the end before method return but the values of the variables are replaced with the
actual values at that point of time.
*/
func funcDefer() {
	i := 1
	fmt.Println("Start ", i)
	i++
	defer fmt.Println("Defer 1 ", i)
	i++
	fmt.Println("Middle ", i)
	i++
	defer fmt.Println("Defer 2 ", i)
	i++
	fmt.Println("End ", i)
	i++
	defer fmt.Println("Defer 3 ", i)
}

func httpCallExample() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()													//here we are calling the Close() before reading the body, but since we are using defer,
	robots, err := ioutil.ReadAll(res.Body)					//this call will be executed before return and hence it'll run just fine.
	if err != nil {																	//The benefit of this is that we can write the code to open and close at one place unlike in other languages
		log.Fatal(err)																//where close needs to be done after the usage is over
	}
	fmt.Printf("%s", robots)
}

/*
Panic is similar to exceptions when go is not able to continue the execution and start panicking.
E.g. the function below panics due to divide by zero error and is not able to then continue the execution hence, not printing the 'End'
*/
func divideByZero() {
	fmt.Println("Start")								//output: Start
	a,b := 1,0
	fmt.Println(a/b)										//panic: runtime error: integer divide by zero
	fmt.Println("End")
}

/*
before the program panics it executes all the defer statements as shown in example below
output :
Start 1
Start 2
Defer statement 2
Defer statement 1
panic: Something bad has happened 2

goroutine 1 [running]:
main.deferAndPanic2()
	/deferPanicRecover.go:93 +0xfe
main.deferAndPanic()
	/deferPanicRecover.go:86 +0xee
main.main()
	/deferPanicRecover.go:17 +0x25
exit status 2*/
func deferAndPanic() {
	fmt.Println("Start 1")
	defer fmt.Println("Defer statement 1")
	deferAndPanic2()
	panic("Something bad has happened 1 ")
	fmt.Println("End 1")
}
func deferAndPanic2() {
	fmt.Println("Start 2")
	defer fmt.Println("Defer statement 2")
	panic("Something bad has happened 2")
	fmt.Println("End 2")
}

/*
output :
Start 1
Start 2
2021/08/15 02:05:55 Recover Functionality: logging error :  Something bad has happened 2
End 1
Defer statement 1
*/
func panicAndRecover() {
	fmt.Println("Start 1")
	defer fmt.Println("Defer statement 1")
	panicAndRecover2()
	//panic("Something bad has happened 1 ")
	fmt.Println("End 1")
}
func panicAndRecover2() {
	fmt.Println("Start 2")
	defer func() {																											// using an anonymous function
		if err := recover(); err != nil {																	//recover() function is called to recover the program from the panic happened in the next line
			log.Println("Recover Functionality: logging error : ", err)
		}
	}()
	panic("Something bad has happened 2")
	fmt.Println("End 2")
}
