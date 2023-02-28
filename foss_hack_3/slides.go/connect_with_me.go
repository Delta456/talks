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
		Type:          "Bold",
		ContentAlign:  "Left",
		Color:         uint(0x00add9),
		TitlePos:      "Top",
		AllowWrapping: true,
	})
	bx.Println("Connect with me 👋",
		fmt.Sprintf(`
    %s: github.com/Delta456
    %s: linkedin.com/in/swastik-baranwal/
    %s: twitter.com/Delta2315`,
			color.HEX("999999", true).Sprint("GitHub"),
			color.HEX("0077b7", true).Sprint("LinkedIn"),
			color.HEX("1da1f2", true).Sprint("Twitter")))

}
