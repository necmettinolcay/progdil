package main

import "fmt"

func main() {
	elements := generateName()
	show(elements)
}

func show(list []string) {
	fmt.Println("***********************************")
	for _, pName := range list {
		fmt.Printf("**  %-26s %3s*\n", pName, "*")
	}
	fmt.Println("***********************************")
}
