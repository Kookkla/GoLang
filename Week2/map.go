package main

import "fmt"

func main() {
	var m map[string]string
	m = make(map[string]string)
	m["nong"] = "AnuchitO"
	fmt.Println(m["nong"])
}
