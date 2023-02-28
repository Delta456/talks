package main

import "github.com/Delta456/box-cli-maker/v2"

func main() {
	bx := box.New(box.Config{
		Px:           2,
		Py:           0,
		Type:         "Single",
		ContentAlign: "Center", // Change this to Right or Center to see the difference
		Color:        "Green",
		TitlePos:     "Bottom",
	})
	bx.Println("FOSS Hack 3.0", "FOSS Hack 3.0 is the third edition of FOSS Hack, \na hackathon to promote Free and Open Source Software by bringing together \nstudents and professionals to build or extend FOSS projects.")
}
