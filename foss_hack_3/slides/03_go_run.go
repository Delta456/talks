package main

import (
	"fmt"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/gookit/color"
)

func main() {
	bx := box.New(box.Config{
		Px:            3,
		Py:            3,
		Type:          "Double Single",
		ContentAlign:  "Left", // Change this to Right or Center to see the difference
		Color:         uint(0x00add9),
		TitlePos:      "Top",
		AllowWrapping: true,
	})
	bx.Println("Installation", fmt.Sprintf("%s install https://github.com/Delta456/box-cli-maker-v2", color.HEX("e5c07b").Sprint("go")))
}
