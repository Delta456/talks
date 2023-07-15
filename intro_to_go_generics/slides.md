# Introduction to Go Generics

This talk aims to deliver about Generics in The Go Programming Language aka [golang.dev](https://go.dev/).

To use generics in Go, Go 1.18 is required to be installed! It is **recommeded** to use Go 1.20.

---

## About me

I am Swastik Baranwal aka `@Delta`:

- A Programmer
- An Open Source Contributor
- A Developer at `@v_language`
- A Volunteer at `@foss_united`
- A Low level, DevOps, Backend enthusiast
- A Technical Writer

---

## What are Generics in Go?

Generics in Go refer to the ability to write `functions`, `data structures`, and `algorithms` that can operate on different types without sacrificing **type safety**. Generics allow developers to write reusable and flexible code that can work with a variety of _data types_.

Unlike other languages Generics in Go have quite a few differences.

---

## The Code

```go
package main

import "fmt"

func main() {
	printStrArr([]string{"Foss", "United", "Delhi"})
	printIntArr([]int{1, 2, 3})
}

func printIntArr(arr []int) {
	fmt.Println(arr)
}

func printStrArr(arr []string) {
	fmt.Println(arr)
}
```

---

## The Code v2

Same as the previous one but this time with **generics**  

```go
package main

import "fmt"

func main() {
	printArr[string]([]string{"Foss", "United", "Delhi"})
	printArr[int]([]int{1, 2, 3})
}

func printArr[T any](t []T) { // any is an alias of interface{}
	fmt.Println(t)
}
```

---

It also works without explictly passing the types.

```go
package main

import "fmt"

func main() {
	printArr([]string{"Foss", "United", "Delhi"})
	printArr([]int{1, 2, 3})
}

func printArr[T any](t []T) {
	fmt.Println(t)
}
```

---

## Multi Type Parameter

Go allows the usage of multi type paramaters i.e. having more than one type parameter.

```go
package main

import "fmt"

func main() {
	printTwoVals(string('A'), rune('A'))

}

func printTwoVals[T, V interface{}](t T, v V) {
	fmt.Println(t, v)
}
```

---

But in this case using `any` for multi type parameter will fail.

```go
package main

import "fmt"

func main() {
	m := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}
	printMap(m)
}

func printMap[K, V any](m map[K]V) {
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
}
```

---

This can be fixed by using `comparable` constraint

```go
package main

import "fmt"

func main() {
	m := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}
	printMap(m)
}

func printMap[K comparable, V any](m map[K]V) {
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
}
```

---

## Constraints

Constraints in Go refer to the ability to apply **restrictions** or **requirements** on the types used with generics. Constraints allow you to define specific conditions that type parameters must satisfy when working with `generic functions`, `structs`, or `interfaces`.

```go
type SignedInt interface {
	int | int8 | int16 | int32 | int64
}
```

---

```go
package main

import "fmt"

type SignedInt interface {
	int | int8 | int16 | int32 | int64
}

func main() {
	signed(uint(67)) // Error: uint does not satisfy SignedInt (uint missing in int | int8 | int16 | int32 | int64)
}

func signed[T SignedInt](t T) {
	fmt.Println(t)
}
```

---

`~` needs to be used in constraint declaration if the type paramater is a type `alias`.

```go
type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
```

The below code would be invalid

```go
type Mysigned int8

type SignedInt interface {
	~Mysigned // INVALID: underlying type of Mysigned is not Mysigned
}
```

---

Constraints can be `combined`ed as well

```go
type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Int interface {
	SignedInt | UnsignedInt
}
```

---

## Generic `stringfy` with `constraints`

```go
func stringify[T fmt.Stringer](slice []T) []string {
	ret := make([]string, 0, len(slice))

	for _, value := range slice {
		ret = append(ret, value.String())
	}

	return ret
}
```

---

```go
type user struct {
	name string
	id   int
}

func (u user) String() string {
	return fmt.Sprintf("{type: \"user\", name: %q, id: %d}", u.name, u.id)
}
```

---

```go
func main() {
	users := []user{
		{name: "Bill", id: 45},
		{name: "Ale", id: 80},
	}
	fmt.Println(stringify(users))
}
```

---

## Type assertion

Type parameters can also be asserted.

```go
func do[T any](i T) {
	switch v := (interface{})(i).(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
}
```

**NOTE**: This is not the _recommended_ way to do this. There is an ongoing proposal to do this in a much better way [`#45380`](https://github.com/golang/go/issues/45380)

---

## Generic Type Aliases

Type aliases can also have `type parameters`.

```go
import "fmt"

type vector[T any] []T

func (v *vector[T]) Push(x T) { *v = append(*v, x) }

func main() {
	var vectorInt vector[int]
	vectorInt.Push(12)
	vectorInt.Push(56)
	vectorInt.Push(45)

	fmt.Println(vectorInt)
}
```

---

## Generic Structs

Structs in Go can also have `type parameters`.

```go
// List is a linked list of values of type T.
type List[T any] struct {
	next *List[T] // this reference to List[T] is OK
	val  T
}

// This type is INVALID.
type P[T1, T2 any] struct {
	F *P[T2, T1] // INVALID; must be [T1, T2]
}
```

---

## When to use type parameters

- Functions that work on `slices`, `maps`, and `channels` of any type.
- General purpose data structures
- When a method looks the same for all types

---

## When not to use type parameters

- When just calling a method on the type argument.

```go
func ReadFour(r io.Reader) ([]byte, error) // GOOD

func ReadFour[T io.Reader](r T) ([]byte, error) // BAD
```

- When the implementation of a common `method` is different for each `method`.
- When the operation is different for each type, even without a `method`.

---

# Miscellaneous 

---

### Generic Variadic Functions

```go
package main

import "fmt"

type Add interface {
	~int | ~int8 | ~int32
}

func printSum[T Add](t ...T) T {
	var sum T
	for _, ele := range t {
		sum += ele
	}
	return sum

}

func main() {

	fmt.Println(printSum(3, 4, 5, 7, 8))
	fmt.Println(printSum([]int32{1, 3, 5, 7, 9, 12}...))
}
```

---

### Empty Set Interface

It is not possible to `embed` constraints in an `interface`.

```go
type SignedInt interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedInt interface {
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// An interface representing an empty type set: there is no type that is both a Signed and UnSigned
type Int interface {
    SignedInt
    UnsignedInt
}
```

---

### Inline constraints

Constaints can also be `inline` and used `anonymously`. It is only recommended when there are only **few type parameters**.

```go
package main

import "fmt"

func plusByFive[T ~int | ~int8 | ~int16 | ~int32 | ~int64](t T) {
	fmt.Println(t + 5)
}

func main() {
	plusByFive(12)
}
```

---

### It took 10 years for Go to have generics

It was due to following factors:

- Complexity would fall on the writer, not on the user
- Wanted to minimize new concepts
- Preserve clarity and simplicity of `Go`
- Writer and user should be able to work independently
- Short build times, fast execution times

---

### Why not use the syntax of C++ and Java Generics?

When parsing code within a function, such as `v := F<T>`, at the point of seeing the `<` it's ambiguous whether we are seeing a `type instantiation` or an `expression` using the `<` operator. This is very difficult to resolve without type information.

Suppose the code:

```go
a, b = w < x, y > (z)
```

Without **type information**, it is impossible to decide whether the right hand side of the assignment is a pair of expressions `(w < x and y > (z))`, or whether it is a generic function instantiation and call that returns two result values `((w<x, y>)(z))`.

It is a key design decision of Go that parsing be possible without type information, which seems impossible when using angle brackets for generics.

---

# Connect with Me

- Twitter: [@Delta2315](https://twitter.com/Delta2315)
- GitHub: [Delta456](https://github.com/Delta456)
- LinkedIn: [Swastik Baranwal](https://www.linkedin.com/in/swastik-baranwal/)
- Dev.to: [Swastik Baranwal](https://dev.to/delta456)

---

# Thank You All!

I hope everyone loved the talk, and I will always appreciated feedback on how I can improve my speaking skills further 😄

The talk is made and presented on [slides](https://github.com/maaslalani/slides).

The code and slides are avaiable on [talks](ttps://github.com/Delta456/talks).
