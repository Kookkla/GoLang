package main

import "fmt"

type Vertex struct {
	Name, Class string
}

var m = map[string]Vertex{
	"Kookkla": Vertex{
		"Tanapong", "Go Lang",
	},
	"Google": Vertex{
		"Google.com", "Go Go Go",
	},
	"P'Nong": Vertex{
		"Nong Nong Nong", "Go Go Go",
	},
}

func main() {
	fmt.Println(m)
}
