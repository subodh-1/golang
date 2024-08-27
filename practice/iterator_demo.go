// Iterator examples
package main

import (
	"fmt"

	"github.com/subodh-1/golang/iter"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// function to add the elements in the List generics
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next

	}
}

// function to iterate the elements of the List Generic
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func main() {
	lst := List[int]{}
	// Add elements in the List
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)
	fmt.Println(lst)

	//Access All elements from the List
	for e := range lst.All() {
		fmt.Println(e)
	}
}
