package main

import "github.com/Delta456/box-cli-maker/v2"

func main() {
	Box := box.New(box.Config{Px: 3, Py: 3, Type: "Single Double", TitlePos: "Top", Color: uint(0x00add9), ContentColor: uint(0xb7b7b7), TitleColor: uint(0x2da44e)})
	Box.Println("Journey 🏎️", `
    - Learnt Go Language to contribute to V Language
    - Wrote basic programs and thought of making a unique project
    - Saw a project in V known as "boxx"
    - Ported that project to Go Language
    - Implemented features and optimizations in the next minor release
    - Began implementing Unicode, emoji, windows console etc
    - The project is matured enough to be used in production after years of development
	`)
}
