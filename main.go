package main

import (
	"fmt"
	"my-plugin/example"
)

func main() {
	fmt.Println("Enter your name")
	var name string
	fmt.Scan(&name)

	ex := &example.HelloRequest{Name: name}
	fmt.Println(ex.ShowMessages())
}