package main

import (
	"fmt"
	"time"
)

func getName() string {
	return "World111!"
}

func main() {
	t := time.Now().Unix()
	s := fmt.Sprintf("%d", t)
	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(s[len(s)-6:])
}

func foo() string {
	return "Hello from go!"
}
