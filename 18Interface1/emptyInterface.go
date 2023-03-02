package main

import "fmt"

func main() {
	found("king")
	found(10)
}

func found(i interface{}) {
	fmt.Printf("type : %T and value: %v\n", i, i)
}
