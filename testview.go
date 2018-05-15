package main

import (
	"image/color"

	"github.com/fogleman/gg"
)

type TestView struct {
	alloc *Allocation
	color color.Color
}

func NewTestView() *TestView {
	return &TestView{
		color: color.RGBA{189, 189, 189, 255},
	}
}

func (c *TestView) Render(context *gg.Context) {
	alloc := c.Allocation()

	context.SetColor(c.color)
	context.DrawRectangle(0, 0, alloc.Width, alloc.Height)
	context.Fill()
}

func (c *TestView) Measure(orient Orientation, forsize float64) (minimum, natural float64) {
	return 20, 100
}

func (c *TestView) Allocate(alloc *Allocation) {
	c.alloc = alloc
}

func (c *TestView) Allocation() Allocation {
	return *c.alloc
}
