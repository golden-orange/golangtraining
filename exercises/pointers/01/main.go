package main

import "fmt"

func main() {
	v := 43   // v contains int 43
	w := v    // w contains int 43
	vw := &v  // vw contains address of v
	ww := &w  // ww contains address of w
	wv := *&v // wv contains int 43
	*vw = 57  // changes v and wv to int 57, but w remains containing int 43

	fmt.Println("'v'\t", v)
	fmt.Println("w := v", w)
	fmt.Println("'vw := &v' assigns to 'vw' the address of 'v':", vw)
	fmt.Println("'ww := &w' assigns to 'ww' the address of 'w':", ww)
	fmt.Printf("'vw' points to the address of v, and is therefore of type %T:", vw)
	fmt.Println("'*vw' points to the value stored 'vw' which is an address containing '43':", *vw)
	fmt.Println("'*ww' points to the value stored in the address of w which is assigned the value of 'v' which is '43':", *ww)
	fmt.Println("\n\n,'wv := *&v': assigns to 'wv' a pointer to the value in the address of v thus returning '43':", wv)
	fmt.Println("\n, '*vw = 57' assigns the value '57' to the address &vw therefore changing the value in the original location 'v'. '*vw' is '57':", *vw, "and 'v' is also '57':", v)
	fmt.Println("\n, However, 'w' still contains the original value of 'v':", w)
	fmt.Println("\n, And 'wv' which contains the value in the address of 'v' (*&v) is changed to '57':", *&v)
}
