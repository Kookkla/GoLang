package main

import "fmt"

func fibonacci() func() int {
	v1 := 0
	v2 := 1
	fmt.Println("v1, v2 : ",v1,v2)

	return func() int{
		// Swap Value
		v:= v1
		v1 = v2  		
		v2 = v + v2

		//fmt.Println("v, v1, v2 : ",v ,v1,v2)

		return v
	}
}

func main() {
	
	fmt.Println("Start")
	f := fibonacci()
	fmt.Println("End")

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// 			V	V1	V2
//				0	1
// Loop1	0	1	1
// Loop2	1	1	2
// Loop3	1	2	3
// Loop4	2	3	5
// Loop5	3	5	8
// Loop6	5	8	13
// Loop7	8	13	21
// Loop8	13	21	34
// Loop9	21	34	55
// Loop10	34	55	89