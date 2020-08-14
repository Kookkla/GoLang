package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	//p = Vertex = &v
	p := &v
	fmt.Printf("%#v\n", v)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%#v\n", p.X)
}
