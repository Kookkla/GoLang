package main

import "fmt"

func show(s []int) {
	fmt.Println(s)
}

func main() {

	var a [3]int
	a[0] = 44
	a[1] = 49

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	show(a)
	show(primes)
}
