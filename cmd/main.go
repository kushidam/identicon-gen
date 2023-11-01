package main

import (
	"flag"
	"os"

	"github.com/nullrocks/identicon"
)


var (
	namespace     = flag.String("namespace", "github", "use namespace")
	blocksSize      = flag.Int("size", 5, "use number of blocks (Size)")
	density    = flag.Int("density", 3, "use density")
	userText    = flag.String("text", "nullrocks", "use decides the resulting figure")
)

func main(){
	flag.Parse()
	// New Generator: Rehuse 
	ig, err := identicon.New(
		*namespace, // Namespace
		*blocksSize,        // Number of blocks (Size)
		*density,        // Density
	)
	
	if err != nil {
		panic(err) // Invalid Size or Density
	}
	
	ii, err := ig.Draw(*userText) // Generate an IdentIcon
	
	if err != nil {
		panic(err) // Text is empty
	}
	
	// File writer
	img, _ := os.Create("icon.png")
	defer img.Close()
	// Takes the size in pixels and any io.Writer
	ii.Png(300, img) // 300px * 300px
}