package main

import "github.com/Delta456/box-cli-maker/v2"

func main() {
	/*
		Box Styles
		---------

		Single
		Double Single
		Double
		Single Double
		Bold
		Round
		Hidden
		Classic

	*/
	Box := box.New(box.Config{Px: 5, Py: 8, Type: "Single", Color: "Red", TitleColor: "HiBlue", ContentColor: "Green"})
	Box.Println("FOSS Hack 3.0", "A hackathon to promote Free and Open Source Software")
}
