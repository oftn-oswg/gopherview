package main

import (
	"bytes"
	"fmt"
	"go/scanner"
	"go/token"
	"image/color"
	"log"
	"strings"

	"github.com/fogleman/gg"
)

type CodeView struct {
	alloc *Allocation

	// Information regarding the source code it contains
	fset        *token.FileSet
	filename    string
	source      []byte
	sourceLines int
	sourceWidth int
	tokens      []CodeViewToken

	// Information regarding the typography
	fontFace   string
	fontSize   float64
	lineHeight float64
}

type CodeViewToken struct {
	pos token.Pos
	tok token.Token
	lit string
}

func (tok *CodeViewToken) Len() int {
	switch tok.tok {
	case token.ADD, token.SUB, token.MUL, token.QUO, token.REM, token.AND,
		token.OR, token.XOR, token.LSS, token.GTR, token.ASSIGN, token.NOT,
		token.LPAREN, token.LBRACK, token.LBRACE, token.COMMA, token.PERIOD,
		token.RPAREN, token.RBRACK, token.RBRACE, token.SEMICOLON, token.COLON:
		// One-character operators
		return 1
	case token.SHL, token.SHR, token.AND_NOT, token.ADD_ASSIGN, token.SUB_ASSIGN,
		token.MUL_ASSIGN, token.QUO_ASSIGN, token.REM_ASSIGN, token.AND_ASSIGN,
		token.OR_ASSIGN, token.XOR_ASSIGN, token.LAND, token.LOR, token.ARROW,
		token.INC, token.DEC, token.EQL, token.NEQ, token.LEQ, token.GEQ,
		token.DEFINE:
		// Two-character operators
		return 2
	case token.SHL_ASSIGN, token.SHR_ASSIGN, token.AND_NOT_ASSIGN, token.ELLIPSIS:
		// Three-character operators
		return 3
	}
	return len(tok.lit)
}

func NewCodeView(filename string, source []byte) *CodeView {
	codeview := &CodeView{
		filename: filename,
		source:   source,

		fontFace:   "/usr/share/fonts/TTF/DejaVuSansMono.ttf",
		fontSize:   16,
		lineHeight: 24,
	}

	// Get line count

	lines := bytes.Split(source, []byte("\n"))
	for line := range lines {
		width := len(lines[line])
		if width > codeview.sourceWidth {
			codeview.sourceWidth = width
		}
		codeview.sourceLines++
	}

	// Tokenize

	codeview.fset = token.NewFileSet()
	scan := scanner.Scanner{}

	file := codeview.fset.AddFile(filename, codeview.fset.Base(), len(source))
	scan.Init(file, source, nil, scanner.ScanComments)

	for {
		pos, tok, lit := scan.Scan()
		if tok == token.EOF {
			break
		}

		codeview.tokens = append(codeview.tokens, CodeViewToken{pos, tok, lit})
	}

	return codeview
}

func (c *CodeView) Render(context *gg.Context) {
	alloc := c.Allocation()

	context.SetColor(PaperColorScheme.Back)
	context.DrawRectangle(0, 0, alloc.Width, alloc.Height)
	context.FillPreserve()
	context.Clip()
	defer context.ResetClip()

	context.SetColor(PaperColorScheme.Text)
	if err := context.LoadFontFace(c.fontFace, c.fontSize); err != nil {
		log.Print(err)
		return
	}

	lineY := c.lineHeight / 2
	lines := bytes.Split(c.source, []byte("\n"))
	currentTokIndex := 0

	for lineIndex, line := range lines {
		// Keep track of the indices to change colors
		type ColorMarker struct {
			Index int
			Color color.Color
		}
		fmt.Printf("\nline %d of %d: %s\n", lineIndex+1, len(lines), line)

		// Add all tokens from this line
		markers := []ColorMarker{{0, PaperColorScheme.Text}}
		for currentTokIndex < len(c.tokens) {
			tok := c.tokens[currentTokIndex]
			pos := c.fset.Position(tok.pos)
			if pos.Line != lineIndex+1 {
				// There are no more tokens on this line
				break
			}
			indexStart := pos.Column - 1
			indexEnd := indexStart + tok.Len()
			if indexEnd > len(line) {
				indexEnd = len(line)
			}
			fmt.Printf("adding token %q %q index %d:%d\n", tok.tok, tok.lit, indexStart, indexEnd)
			markers = append(markers,
				ColorMarker{indexStart, PaperColorScheme.Tokens[tok.tok]},
				ColorMarker{indexEnd, PaperColorScheme.Text})
			currentTokIndex++
		}
		markers = append(markers, ColorMarker{len(line), PaperColorScheme.Text})

		// Iterate each consecutive pair of markers that has text in between
		var lineX float64
		for index := 0; index < len(markers)-1; index++ {
			markerA, markerB := markers[index], markers[index+1]
			if markerA.Index == markerB.Index {
				continue
			}
			str := string(line[markerA.Index:markerB.Index])
			str = strings.Replace(str, "\t", "    ", -1)

			width, _ := context.MeasureString(str)
			context.SetColor(markerA.Color)
			context.DrawStringAnchored(str, lineX, lineY, 0, 0.5)

			lineX += width
		}
		lineY += c.lineHeight
	}
	context.Fill()
}

func (c *CodeView) Measure(orient Orientation, forsize float64) (minimum, natural float64) {
	switch orient {
	case Horizontal:
		// Natural and minimum size is the width of the longest line
		minimum = float64(c.sourceWidth) * c.fontSize
		natural = minimum

	case Vertical:
		// Natural size is the number of lines times the line height
		// Minimum size is 3 lines
		minimum = 3 * c.lineHeight
		natural = float64(c.sourceLines) * c.lineHeight
	}
	return
}

func (c *CodeView) Allocate(alloc *Allocation) {
	c.alloc = alloc
}

func (c *CodeView) Allocation() Allocation {
	return *c.alloc
}
