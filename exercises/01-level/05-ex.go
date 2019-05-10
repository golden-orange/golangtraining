package main
import "fmt"

type beer int

var x beer
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}


// now use CONVERSION to convert the TYPE of the VALUE stored in “x”
// to the UNDERLYING TYPE
// 1. then use the “=” operator to ASSIGN that value to “y”
// 2. print out the value stored in “y”