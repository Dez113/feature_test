package main

import "fmt"


func main(){
	someVariable := "TESTING!"
	fmt.Println("Hello_world!", someVariable)
	fmt.Println("Hello_world!", someVariable)
	fmt.Println("Hello_world!", someVariable)
	fmt.Println("Hello_world!", someVariable)
	fmt.Println(foo())
	return
}

func foo() string {
	return "Hello from go!"
}
