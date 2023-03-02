// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {

// 	fmt.Println("Addition of two numbers")
// 	fmt.Println("Enter the value of num1 and num2")

// 	reader1 := bufio.NewReader(os.Stdin)
// 	reader2 := bufio.NewReader(os.Stdin)

// 	num1, _ := reader1.ReadString('\n')
// 	num2, _ := reader2.ReadString('\n')

// 	x, err := strconv.ParseFloat(strings.TrimSpace(num1), 64)
// 	y, err := strconv.ParseFloat(strings.TrimSpace(num2), 64)

// 	// var z float32

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("sum of two number is : ", x+y)
// 		fmt.Printf("sum of two number is : %.3f", x+y)
// 	}
// }
