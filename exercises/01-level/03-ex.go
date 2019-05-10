package main
import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%T\t%v\t%T\t", x, y, z)

	fmt.Println(s)

}



// Using the code from the previous exercise,
// 1. At the package level scope, assign the following values to the three variables
// a. for x assign 42
// b. for y assign “James Bond”
// c. for z assign true
// 2. in func main
// a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the
// returned value of TYPE string using the short declaration operator to a
// VARIABLE with the IDENTIFIER “s”
// b. print out the value stored by variable “s”
// code: here’s the solution: https://play.golang.org/p/QFctSQB_