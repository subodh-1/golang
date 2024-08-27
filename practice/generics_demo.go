/*
*
Generics Demo
*/
package main

import (
	"fmt"
)

type Pack[U any, T any] struct {
	First  U
	Second T
}

func main() {
	f1 := Pack[int, string]{First: 1, Second: "Second is string"}
	f2 := Pack[string, string]{First: "First is string", Second: "Second is string"}
	arr :=
		[5]string{"first", "second", "third", "fourth", "fifth"}
	f3 := Pack[int, [5]string]{First: 1, Second: arr}
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f3.Second)    // Accessing second element
	fmt.Println(f3.Second[3]) // Accessing array item
}
