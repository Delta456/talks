package main

import "fmt"

func stringify[T fmt.Stringer](slice []T) []string {
	ret := make([]string, 0, len(slice))

	for _, value := range slice {
		ret = append(ret, value.String())
	}

	return ret
}

type user struct {
	name string
	id   int
}

func (u user) String() string {
	return fmt.Sprintf("{type: \"user\", name: %q, id: %d}", u.name, u.id)
}

func main() {
	users := []user{
		{name: "Rahul", id: 45},
		{name: "Ram", id: 80},
	}
	fmt.Println(stringify(users))
}
