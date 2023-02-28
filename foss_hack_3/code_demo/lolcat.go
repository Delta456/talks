package main

import (
	"math"
	"strings"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/gookit/color"
)

func main() {
	// Know more about lolcat: https://www.tecmint.com/lolcat-command-to-output-rainbow-of-colors-in-linux-terminal/
	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan"})
	Box.Println(lolcat("FOSS Hack"), lolcat("A hackathon to promote Free and Open Source Software"))
}

func lolcat(str string) string {
	var output string
	freq := float64(0.1)
	for _, s := range strings.Split(str, "") {
		output += normalStyle(freq, s)
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
