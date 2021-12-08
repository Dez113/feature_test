package main

import "fmt"
func getName() string{
	return "World111!"
}

func main(){
	someVariable := "TESTING!"
	fmt.Printfghln("Hello_world!", someVariable)
	fmt.Println("Hello_world!", someVariable)
	fmt.Println(foo())
	return
}

func foo() string {
	return "Hello from go!"
}
