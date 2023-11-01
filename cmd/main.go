package main

import (
	"os"

	"github.com/nullrocks/identicon"
)

func main(){
	
	// New Generator: Rehuse 
	ig, err := identicon.New(
		"github", // Namespace
		5,        // Number of blocks (Size)
		3,        // Density
	)
	
	if err != nil {
		panic(err) // Invalid Size or Density
	}
	
	username := "nullrocks"      // Text - decides the resulting figure
	ii, err := ig.Draw(username) // Generate an IdentIcon
	
	if err != nil {
		panic(err) // Text is empty
	}
	
	// File writer
	img, _ := os.Create("icon.png")
	defer img.Close()
	// Takes the size in pixels and any io.Writer
	ii.Png(300, img) // 300px * 300px
}