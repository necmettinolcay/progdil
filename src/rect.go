package main

import "fmt"

type Rectangle struct {
	length, width int
	name          string
}

func main() {
	pr := new(Rectangle)
	pr.width = 6
	pr.length = 8
	pr.name = "ptr_to_rectangle"
	fmt.Println("Rectangle pr as address is: ", pr)
	fmt.Println("Rectangle pr as value is: ", *pr)
}
