package main

import "fmt"

// the following should print:
//0
//43
//0
// func main() {
// 	x := 0
// 	foo(x)
// 	fmt.Println(x)
// }
//
// func foo(y int) {
// 	fmt.Println(y)
// 	y = 43
// 	fmt.Println(y)
// }

// the following should print:
// 0
// 0
// 43
// 43
// 43 because *y changed the value in x
// what actually printed was:
// address
// address
// address
// address
// 43
// func main() {
// 	x := 0
// 	foo(&x)
// 	fmt.Println(x)
// }
//
// func foo(y *int) {
// 	fmt.Println(y)
// 	fmt.Println(y)
// 	*y = 43
// 	fmt.Println(y)
// 	fmt.Println(y)
// }

func main() {
	x := 0
	foo(&x)
	fmt.Println(x) // outputs '99'
}

func foo(y *int) {
	fmt.Println(y)  // outputs address (&x)
	fmt.Println(*y) // ouputs '0'
	*y = 99
	fmt.Println(y)  // outputs address (&x)
	fmt.Println(*y) // outputs '99'
}
