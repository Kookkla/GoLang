package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("%T=>%v\n", primes, primes)

	var s []int = primes[0:6]
	fmt.Printf("%T=>%v\n", s, s)

	s[0] = 77

	fmt.Printf("%T=>%v\n", primes, primes)
	fmt.Printf("%T=>%v\n", s, s)

	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
