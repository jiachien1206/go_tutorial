package main

import "fmt"

func main() {
	a := [3]string{"a", "b"}
	fmt.Println(a)
	fmt.Printf("%#v\n", a[2])
	fmt.Printf("Type: %T\n",a[2])
	fmt.Printf("Type: %T\n",a)


	b := [3]int{1,2}
	fmt.Println(b)
	fmt.Printf("%#v\n", b[2])
	
	c := make([]byte, 5, 5)
	fmt.Println(c)

	d := []byte{104, 101, 108}
	fmt.Println(d)
	str := string(d)
	fmt.Println(str)

	fmt.Println(string(126)) // ASCII
}

