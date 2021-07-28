package main

import "fmt"


func main(){
	var someVariable = "disfjldkfjsdf"
	fmt.Println("Hello_world!", someVariable)
	fmt.Println(foo())
	return
}

func foo() string {
	return "Hello from go!"
}
