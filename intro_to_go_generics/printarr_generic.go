package main

import "fmt"

func main() {
	printArr([]string{"Foss", "United", "Delhi"})
	printArr([]int{1, 2, 3})
}

func printArr[T any](t []T) {
	fmt.Println(t)
}
