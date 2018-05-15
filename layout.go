package main

import "github.com/fogleman/gg"

type Layout struct {
	View

	Children []View
	alloc    *Allocation
}

func (c *Layout) Add(child View) {
	c.Children = append(c.Children, child)
}

func (c *Layout) Remove(child View) {
	list := c.Children[:0]
	for _, item := range c.Children {
		if item != child {
			list = append(list, item)
		}
	}
	c.Children = list
}

func (c *Layout) Render(context *gg.Context) {
	context.Push()
	defer context.Pop()

	// FIXME: https://github.com/fogleman/gg/issues/27
	// alloc := c.Allocation()
	// context.DrawRectangle(alloc.X, alloc.Y, alloc.Width, alloc.Height)
	// context.Clip()
	// defer context.ResetClip()

	for _, child := range c.Children {
		func() {
			context.Push()
			defer context.Pop()

			alloc := child.Allocation()
			context.Translate(alloc.X, alloc.Y)
			child.Render(context)
		}()
	}
}

func (c *Layout) Allocation() Allocation {
	return *c.alloc
}

func (c *Layout) Len() int {
	return len(c.Children)
}
