package main

import "fmt"

type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Int interface {
	SignedInt
	UnsignedInt
}

func works[T Int](t T) {
	fmt.Println(t)
}

func main() {
	works(12)
}
