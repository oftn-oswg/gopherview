package main

import "math"

type Scroller interface {
	ScrollTo(offset float64)
	ScrollOffset() float64
	ScrollSize() float64
}

type LinearLayout struct {
	Layout

	Orientation Orientation
	Spacing     float64
	Margin      float64

	scrollOffset float64
	scrollSize   float64
}

func NewLinearLayout(orient Orientation, spacing, margin float64) *LinearLayout {
	return &LinearLayout{
		Orientation: orient,
		Margin:      margin,
		Spacing:     spacing,
	}
}

func (s *LinearLayout) Measure(orient Orientation, forsize float64) (minimum, natural float64) {
	// Compute minimum and natural size if allocated a height of `forsize`.
	// We can accomplish this by adding up the minimum and natural
	// widths/heights of all the children along with the total spacing
	// between them.
	if orient == s.Orientation {
		// We are measuring the width of a horizontal scroller.
		// Find the sum of all minimum and natural measures of all children
		// and add margins.
		minimum = 2 * s.Margin
		natural = minimum + s.Spacing*float64(s.Len()-1)

		for _, child := range s.Layout.Children {
			minimumChild, naturalChild := child.Measure(orient, forsize)
			minimum += minimumChild
			natural += naturalChild
		}
	} else {
		// We are measuring the height of a horizontal scroller.
		// Find the largest minimum and natural measure for all children
		// and add margins.
		minimum = 2 * s.Margin
		natural = minimum

		var minimumChildMax, naturalChildMax float64
		for _, child := range s.Layout.Children {
			minimumChild, naturalChild := child.Measure(orient, forsize)
			minimumChildMax = math.Max(minimumChildMax, minimumChild)
			naturalChildMax = math.Max(naturalChildMax, naturalChild)
		}

		minimum += minimumChildMax
		natural += naturalChildMax
	}

	return
}

func (s *LinearLayout) Allocate(alloc *Allocation) {
	s.Layout.alloc = alloc

	var forsize float64
	var space float64
	switch s.Orientation {
	case Horizontal:
		forsize, space = alloc.Height, alloc.Width
	case Vertical:
		forsize, space = alloc.Width, alloc.Height
	}

	numberOfChildren := len(s.Layout.Children)
	if numberOfChildren == 0 {
		return
	}

	// Subtract margin
	forsize -= 2 * s.Margin
	space -= 2 * s.Margin

	// Calculate distribution
	var sizesSum float64
	sizes := make([]float64, numberOfChildren)
	for i, child := range s.Layout.Children {
		_, natural := child.Measure(s.Orientation, forsize)
		sizes[i] = natural
		sizesSum += natural
	}
	sizeTotal := sizesSum + float64(numberOfChildren-1)*s.Spacing
	sizeExtra := space - sizeTotal
	if sizeExtra < 0 {
		sizeExtra = 0
	}
	toDistribute := sizeExtra / float64(numberOfChildren)

	// Set allocation to all of child nodes
	pointer := s.Margin
	for i, child := range s.Layout.Children {
		natural := sizes[i] + toDistribute

		childAlloc := &Allocation{
			X:      s.scrollOffset + pointer,
			Y:      s.Margin,
			Width:  natural,
			Height: forsize,
		}

		if s.Orientation == Vertical {
			childAlloc.X, childAlloc.Y = childAlloc.Y, childAlloc.X
			childAlloc.Width, childAlloc.Height = childAlloc.Height, childAlloc.Width
		}

		child.Allocate(childAlloc)

		pointer += natural + s.Spacing
	}

	s.scrollSize = sizeTotal
}

// Scroller interface

func (s *LinearLayout) ScrollTo(offset float64) {
	s.scrollOffset = offset
}

func (s *LinearLayout) ScrollOffset() float64 {
	return s.scrollOffset
}

func (s *LinearLayout) ScrollSize() float64 {
	return s.scrollSize
}
