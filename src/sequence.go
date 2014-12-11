package main

import "fmt"

type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}

func main() {
	seq := make(Sequence, 0)
	seq = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(seq.Len())
	fmt.Println(seq)
}
