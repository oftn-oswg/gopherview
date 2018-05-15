package main

import (
	"github.com/fogleman/gg"
)

type Orientation int

const (
	Horizontal Orientation = iota
	Vertical
)

// Allocation is a rectangular region which has been allocated to the element
// by its parent. It is a subregion of its parents allocation.
type Allocation struct {
	X, Y, Width, Height float64
}

type Measurer interface {
	Measure(orient Orientation, forsize float64) (minimum, natural float64)
}

type Allocator interface {
	Allocate(alloc *Allocation)
	Allocation() Allocation
}

type Renderer interface {
	Render(context *gg.Context)
}

type View interface {
	Measurer
	Allocator
	Renderer
}
