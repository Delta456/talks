package main

import (
	"fmt"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/gookit/color"
)

const Art = ` 
.-------------.  ================================================================================================
/  .           /  ██████   ██████  ██   ██      ██████ ██      ██     ███    ███  █████  ██   ██ ███████ ██████  
@  =.          +  ██   ██ ██    ██  ██ ██      ██      ██      ██     ████  ████ ██   ██ ██  ██  ██      ██   ██ 
@    :         +  ██████  ██    ██   ███       ██      ██      ██     ██ ████ ██ ███████ █████   █████   ██████  
@              +  ██   ██ ██    ██  ██ ██      ██      ██      ██     ██  ██  ██ ██   ██ ██  ██  ██      ██   ██ 
@              +  ██████   ██████  ██   ██      ██████ ███████ ██     ██      ██ ██   ██ ██   ██ ███████ ██   ██ 
./############./  ================================================================================================
`

func main() {
	Box := box.New(box.Config{Px: 3, Py: 3, Type: "Single Double", Color: uint(0x00add9), ContentColor: uint(0x01aeb0)})
	// All ANSI Art may not fit inside the box as they can exceed the terminal width
	Box.Println("", Art)
	fmt.Println(fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t\t\t%s\n", color.HEX("00c0b0").Sprint("- Made by Swastik Baranwal")))
}
