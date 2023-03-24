package main

import (
	"fmt"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/gookit/color"
)

func main() {
	Box := box.New(box.Config{Px: 3, Py: 3, Type: "Double", TitlePos: "Top", Color: uint(0x00add9), AllowWrapping: true})
	Box.Println(color.HEX("ffd700").Sprint("Wins Along The Way 🥇"), fmt.Sprintf(`
    Used in %s project by %s.
    Featured on %s Radar, %s India, Awesome %s, %s Weekly etc
	`, color.HEX("326de6").Sprint("minikube"),
		color.HEX("2e6ce6").Sprint("Kubernetes"),
		color.HEX("2fbb4f").Sprint("GitHub"),
		color.HEX("2fbb4f").Sprint("GitHub"),
		color.HEX("00add8").Sprint("Go"),
		color.HEX("00add8").Sprint("Go")))
}
