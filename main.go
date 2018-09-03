package main

import "fmt"

func main() {
	fmt.Println("hello go mod")
}

func Hello(name string) string {
	s := "from hello" + name
	return s
}
