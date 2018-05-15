package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"

	"github.com/fogleman/gg"
)

// This is a comment.
func main() {
	overflowB := NewLinearLayout(Horizontal, 10, 10)
	overflowC := NewLinearLayout(Vertical, 10, 0)

	overflowC.Add(NewTestView())
	overflowC.Add(NewTestView())
	overflowC.Add(NewTestView())

	filename := "main.go"
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	overflowB.Add(NewTestView())
	overflowB.Add(NewCodeView(filename, source))
	overflowB.Add(overflowC)

	alloc := &Allocation{
		X:      0,
		Y:      0,
		Width:  1920,
		Height: 1080,
	}

	overflowB.Allocate(alloc)

	context := gg.NewContext(int(alloc.Width), int(alloc.Height))
	context.SetColor(color.RGBA{255, 255, 255, 255})
	context.Clear()

	overflowB.Render(context)

	out := "out.png"
	if len(os.Args) >= 2 {
		out = os.Args[1]
	}

	fmt.Printf("Saving %dx%d image in %s\n", context.Width(), context.Height(), out)

	if err := context.SavePNG(out); err != nil {
		log.Fatal(err)
	}
}
