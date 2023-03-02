package main

import "fmt"

func main() {
	fmt.Println("Welcome to the fuction learning in golang")
	greetings()
	result := adder(5, 9)
	fmt.Println("value is :", result)

	proRes, myMessage := proAdder(1, 2, 3, 4)
	fmt.Println("Pro adder values :", proRes, myMessage)
}

func greetings() {
	fmt.Println("Hi Good evening!!")
}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi pro adder returning"
}
