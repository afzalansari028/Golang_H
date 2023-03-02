// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	fmt.Println("Enter your number to find factorial")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	val, err := strconv.ParseInt(strings.TrimSpace(input), 0, 64)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Factorial is :", factorial(val))
// }

// //factorial function using recursion
// func factorial(n int64) int64 {

// 	if n == 1 {
// 		return 1
// 	}
// 	return n * factorial(n-1)
// }
