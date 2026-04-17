package main

import (
	"github.com/renatopp/x/dsx"
)

type List[T any] []T

func main() {
	heap := dsx.NewHeap[string]()

	heap.Push(1, "apple")
	heap.Push(1, "banana")
	heap.Push(1, "pineapple")

	println(heap.Get(0))
	println(heap.Get(1))
	println(heap.Get(2))
}
