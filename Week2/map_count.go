package main
import "fmt"

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck
	w := WordCount(s)
	fmt.Println(w)
}