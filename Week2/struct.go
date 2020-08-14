package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	a := Vertex{3, 3}
	fmt.Println(a)
	fmt.Println(a.X)
	a.X = 77
	fmt.Println(a)
}
