package main

import (
	"github.com/Delta456/box-cli-maker/v2"
)

func main() {
	// Know more about ANSI Escape Codes: https://en.wikipedia.org/wiki/ANSI_escape_code
	Box := box.New(box.Config{Px: 2, Py: 1, Type: "Single", Color: "Blue", TitlePos: "Top"})
	Box.Println("\033[1mFOSS\033[0m, Hack", "Btw \033[1mANSI Text works here too\033[0m, very nice")
}
