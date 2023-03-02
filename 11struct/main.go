package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct in golang")
	//no inheritance in golang; No super parent

	afzal := User{"Afzal", "afzal@gmail.com", true, 23}
	fmt.Println(afzal)
	fmt.Printf("Afzal details are : %+v\n", afzal)
	fmt.Printf("name is %v and email is %v\n", afzal.Name, afzal.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
