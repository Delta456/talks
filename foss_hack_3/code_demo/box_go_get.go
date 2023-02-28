package main

import (
	"fmt"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/gookit/color"
)

func main() {
	Box := box.New(box.Config{Px: 3, Py: 2, Type: "Round", Color: "Yellow", AllowWrapping: true})
	Box.Println("", fmt.Sprintf("%s Update available %s -> %s\nRun %s then %s in your project directory", color.Red.Sprintf("Box CLI Maker"), color.Magenta.Sprint("v2.0.0"), color.Green.Sprint("v2.3.0"), color.Cyan.Sprintf("go install -v -u github.com/Delta456/box-cli-maker/v2@master"), color.Cyan.Sprint("go mod tidy")))
}
