package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/gookit/color"
)

func main() {
	Box := box.New(box.Config{Px: 3, Py: 3, Type: "Single", Color: uint(0x00add9), TitlePos: "Top"})
	Box.Println(color.HEX("ffd700").Sprint("Features 🏆"), fmt.Sprintf(`
    Create CLI %s 📦 in upto 8 %s
	Also make your own box style
    Support for %s 🎨 and %s
    Make %s, %s, %s boxes 📋
    %s な, %s 💖 and %s Support
    Text %s and %s Title Positions 📐`,
		color.BgLightGreen.Sprint("Boxes"),
		color.New(color.OpUnderscore).Sprint("inbuilt styles"),
		color.Yellow.Sprint("True Color"),
		color.New(color.OpFuzzy).Sprint("Terminal Color Detection"),
		lolcat("ANSI", 0),
		lolcat("Tabbed", 1),
		lolcat("Multi-Line", 2),
		color.HEX("007acc").Sprint("UTF-8"),
		color.HEX("fcd252").Sprint("Emoji"),
		color.HEX("5e4a80").Sprint("Windows Console"),
		color.New(color.OpFastBlink).Sprint("Alignment"),
		color.New(color.OpItalic).Sprint("Custom")))
}

func lolcat(str string, choice int) string {
	var output string
	freq := float64(0.1)
	for _, s := range strings.Split(str, "") {
		if choice == 1 {
			output += normalStyle(freq, s)
		} else if choice == 2 {
			output += greenStyle(freq, s)
		} else {
			output += redStyle(freq, s)
		}
		freq += 0.1

	}
	return output
}

func normalStyle(num float64, s string) string {
	red := uint8(math.Sin(num+0)*127 + 128)
	green := uint8(math.Sin(num+2)*127 + 128)
	blue := uint8(math.Sin(num+4)*127 + 128)

	return color.Rgb(red, blue, green).Sprint(s)
}

func redStyle(num float64, s string) string {
	red := uint8(math.Sin(num+0)*127 + 128)

	return color.Rgb(red, 0, 0).Sprint(s)
}

func greenStyle(num float64, s string) string {
	green := uint8(math.Sin(num+2)*127 + 128)

	return color.Rgb(0, green, 0).Sprint(s)
}
