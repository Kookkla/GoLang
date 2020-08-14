package main

import "fmt"

type Student struct {
	Name, Class string
}

var m = map[string]Student{
	"Kookkla": Student{
		"Tanapong", "ป.1/1",
	},
	"Mameaw": Student{
		"Jirapong", "ป.1/2",
	},
}

func main() {
	fmt.Println(m)
}
