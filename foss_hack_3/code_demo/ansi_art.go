package main

import "github.com/Delta456/box-cli-maker/v2"

func main() {
	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan", ContentColor: "Green"})
	// All ANSI Art may not fit inside the box as they can exceed the terminal width
	Box.Println("", `
	
___________________    _________ _________   ___ ___                __    
\_   _____/\_____  \  /   _____//   _____/  /   |   \_____    ____ |  | __
 |    __)   /   |   \ \_____  \ \_____  \  /    ~    \__  \ _/ ___\|  |/ /
 |     \   /    |    \/        \/        \ \    Y    // __ \\  \___|    < 
 \___  /   \_______  /_______  /_______  /  \___|_  /(____  /\___  >__|_ \
     \/            \/        \/        \/         \/      \/     \/     \/
`)
}
