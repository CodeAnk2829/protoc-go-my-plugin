package main

import (
	"fmt"
	"my-plugin/example"
)

func main() {
	ex := &example.HelloResponse{Message: "Hi Ankit"}
	fmt.Println(ex.ShowMessages())
}