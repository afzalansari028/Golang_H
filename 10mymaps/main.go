package main

import "fmt"

func main() {
	fmt.Println("Welcome to the learning of Map in GO")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of languages is ", languages)

	fmt.Println("JS is shorts for :", languages["JS"])
	//	fmt.Println("PY is shorts for :", languages["PY"])

	delete(languages, "RB")
	fmt.Println("List of languages is ", languages)

	//loops are interesting in golang

	for key, value := range languages { //traversing the map
		fmt.Printf("For key %v and for value %v\n", key, value)
	}

}
