package main

import (
	"go/token"
	"image/color"
)

type ColorScheme struct {
	Text, Back                         color.Color
	SelectText, SelectBack             color.Color
	GutterText, GutterBack             color.Color
	GutterSelectText, GutterSelectBack color.Color
	Tokens                             map[token.Token]color.Color
}

var base00 = color.RGBA{0xf3, 0xf3, 0xf3, 0xff}
var base01 = color.RGBA{0xd0, 0xd0, 0xd0, 0xff}
var base02 = color.RGBA{0xb0, 0xb0, 0xb0, 0x21}
var base03 = color.RGBA{0x94, 0x94, 0x94, 0xff}
var base04 = color.RGBA{0x44, 0x44, 0x44, 0xff}
var base05 = color.RGBA{0x4d, 0x4d, 0x4c, 0xff}
var base06 = color.RGBA{0x26, 0x26, 0x26, 0xff}
var base07 = color.RGBA{0xc5, 0xc8, 0xc6, 0xff}
var base08 = color.RGBA{0x89, 0x59, 0xa8, 0xff}
var base09 = color.RGBA{0xd7, 0x5f, 0x00, 0xff}
var base0A = color.RGBA{0x42, 0x71, 0xae, 0xff}
var base0B = color.RGBA{0x71, 0x8c, 0x00, 0xff}
var base0C = color.RGBA{0x3e, 0x99, 0x9f, 0xff}
var base0D = color.RGBA{0x00, 0x5f, 0x87, 0xff}
var base0E = color.RGBA{0xd7, 0x00, 0x5f, 0xff}
var base0F = color.RGBA{0xdf, 0x00, 0x00, 0xff}

var PaperColorScheme = ColorScheme{
	Text:             base05,
	Back:             base00,
	SelectText:       base05,
	SelectBack:       base02,
	GutterText:       base05,
	GutterBack:       base00,
	GutterSelectText: base03,
	GutterSelectBack: base00,
	Tokens: map[token.Token]color.Color{
		token.ILLEGAL: base05,
		token.EOF:     base05,
		token.COMMENT: base03,

		// Identifiers and basic type literals
		// (these tokens stand for classes of literals)
		token.IDENT:  base08, // main
		token.INT:    base09, // 12345
		token.FLOAT:  base09, // 123.45
		token.IMAG:   base09, // 123.45i
		token.CHAR:   base09, // 'a'
		token.STRING: base0B, // "abc"

		// Operators and delimiters
		token.ADD: base05, // +
		token.SUB: base05, // -
		token.MUL: base05, // *
		token.QUO: base05, // /
		token.REM: base05, // %

		token.AND:     base05, // &
		token.OR:      base05, // |
		token.XOR:     base05, // ^
		token.SHL:     base05, // <<
		token.SHR:     base05, // >>
		token.AND_NOT: base05, // &^

		token.ADD_ASSIGN: base05, // +=
		token.SUB_ASSIGN: base05, // -=
		token.MUL_ASSIGN: base05, // *=
		token.QUO_ASSIGN: base05, // /=
		token.REM_ASSIGN: base05, // %=

		token.AND_ASSIGN:     base05, // &=
		token.OR_ASSIGN:      base05, // |=
		token.XOR_ASSIGN:     base05, // ^=
		token.SHL_ASSIGN:     base05, // <<=
		token.SHR_ASSIGN:     base05, // >>=
		token.AND_NOT_ASSIGN: base05, // &^=

		token.LAND:  base05, // &&
		token.LOR:   base05, // ||
		token.ARROW: base05, // <-
		token.INC:   base05, // ++
		token.DEC:   base05, // --

		token.EQL:    base05, // ==
		token.LSS:    base05, // <
		token.GTR:    base05, // >
		token.ASSIGN: base05, // =
		token.NOT:    base05, // !

		token.NEQ:      base05, // !=
		token.LEQ:      base05, // <=
		token.GEQ:      base05, // >=
		token.DEFINE:   base05, // :=
		token.ELLIPSIS: base05, // ...

		token.LPAREN: base05, // (
		token.LBRACK: base05, // [
		token.LBRACE: base05, // {
		token.COMMA:  base05, // ,
		token.PERIOD: base05, // .

		token.RPAREN:    base05, // )
		token.RBRACK:    base05, // ]
		token.RBRACE:    base05, // }
		token.SEMICOLON: base05, // ;
		token.COLON:     base05, // :

		// Keywords
		token.BREAK:    base0E,
		token.CASE:     base0E,
		token.CHAN:     base0E,
		token.CONST:    base0E,
		token.CONTINUE: base0E,

		token.DEFAULT:     base0E,
		token.DEFER:       base0E,
		token.ELSE:        base0E,
		token.FALLTHROUGH: base0E,
		token.FOR:         base0E,

		token.FUNC:   base0E,
		token.GO:     base0E,
		token.GOTO:   base0E,
		token.IF:     base0E,
		token.IMPORT: base0E,

		token.INTERFACE: base0E,
		token.MAP:       base0E,
		token.PACKAGE:   base0E,
		token.RANGE:     base0E,
		token.RETURN:    base0E,

		token.SELECT: base0E,
		token.STRUCT: base0E,
		token.SWITCH: base0E,
		token.TYPE:   base0E,
		token.VAR:    base0E,
	},
}
