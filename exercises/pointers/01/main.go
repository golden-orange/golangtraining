package main

import "fmt"

func main() {
	v := 43
	w := v
	vw := &v
	ww := &w
	wv := *&v
	*vw = 57

	fmt.Println("'v'\t", v)
	fmt.Println("w := v", w)
	fmt.Println("'vw := &v' assigns to 'vw' the address of 'v':", vw)
	fmt.Println("'ww := &w' assigns to 'ww' the address of 'w':", ww)
	fmt.Printf("'vw' points to the address of v, and is therefore of type %T:", vw)
	fmt.Println("'*vw' points to the value stored 'vw' which is an address containing '43':", *vw)
	fmt.Println("'*ww' points to the value stored in the address of w which is assigned the value of 'v' which is '43':", *ww)
	fmt.Println("\n\n,'wv := *&v': assigns to 'wv' a pointer to the value in the address of v thus returning '43':", wv)
	fmt.Println("\n\n, '*vw = 57' assigns the value '57' to the address &vw therefore changing the value in the original location 'v'. '*vw' is '57':", *vw, "and 'v' is also '57':", v)
	av := &v
	fmt.Println("'& points to the address of a variable. Therefore:'")
	fmt.Println("'av := &v'", av)

	pv := *av
	fmt.Println("'*' points to the value in the address. Therefore:")
	fmt.Println("'pv := *av'", pv)
}
