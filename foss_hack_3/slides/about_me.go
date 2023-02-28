package main

import (
	"fmt"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/gookit/color"
)

func main() {
	Box := box.New(box.Config{Px: 3, Py: 3, Type: "Round", Color: uint(0x00add9), TitlePos: "Top"})
	Box.Println(color.HEX("dd081a").Sprint("Who am I 🤔"), fmt.Sprintf(`
    %s 📝 & %s Contributor
    %s at The %s Programming Language
    Writes in %s, %s, %s, %s and %s
    Indulges in %s & %s 🥰
    Tinkers with %s, %s 💻 and lives in the %s`,
		color.HEX("737373").Sprint("Technical Writer"),
		color.HEX("3da751").Sprint("Open Source"),
		color.New(color.OpBold).Sprint("Core Developer"),
		color.HEX("4f87c4").Sprint("V"), color.HEX("f34b7d").Sprint("C++"),
		color.HEX("555555").Sprint("C"), color.HEX("00add8").Sprint("Go"),
		color.HEX("3572a5").Sprint("Python"), color.HEX("4f87c4").Sprint("V"),
		color.HEX("194992").Sprint("Anime"), color.HEX("61ad98").Sprint("Manga"),
		color.HEX("ffd200").Sprint("compilers"),
		color.HEX("dd4814").Sprint("operating systems"),
		color.HEX("ffbd63").Sprint("terminal")))
}
