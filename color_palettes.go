package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/sourcegraph/syntaxhighlight"
)

var (
	LightColorPalettes = ColorPalettes{
		syntaxhighlight.String:        termbox.ColorYellow,
		syntaxhighlight.Keyword:       termbox.ColorBlue,
		syntaxhighlight.Comment:       termbox.Attribute(243),
		syntaxhighlight.Type:          termbox.ColorMagenta,
		syntaxhighlight.Literal:       termbox.ColorMagenta,
		syntaxhighlight.Punctuation:   termbox.ColorRed,
		syntaxhighlight.Plaintext:     termbox.ColorRed,
		syntaxhighlight.Tag:           termbox.ColorBlue,
		syntaxhighlight.HTMLTag:       termbox.ColorGreen,
		syntaxhighlight.HTMLAttrName:  termbox.ColorBlue,
		syntaxhighlight.HTMLAttrValue: termbox.ColorGreen,
		syntaxhighlight.Decimal:       termbox.ColorBlue,
	}

	DarkColorPalettes = ColorPalettes{
		syntaxhighlight.String:        termbox.Attribute(20),
		syntaxhighlight.Keyword:       termbox.ColorBlue,
		syntaxhighlight.Comment:       termbox.Attribute(200),
		syntaxhighlight.Type:          termbox.Attribute(2000),
		syntaxhighlight.Literal:       termbox.Attribute(2000),
		syntaxhighlight.Punctuation:   termbox.ColorRed,
		syntaxhighlight.Plaintext:     termbox.ColorBlue,
		syntaxhighlight.Tag:           termbox.ColorBlue,
		syntaxhighlight.HTMLTag:       termbox.Attribute(20000),
		syntaxhighlight.HTMLAttrName:  termbox.ColorBlue,
		syntaxhighlight.HTMLAttrValue: termbox.ColorGreen,
		syntaxhighlight.Decimal:       termbox.ColorBlue,
	}
)

type ColorPalettes map[syntaxhighlight.Kind]termbox.Attribute

func (c ColorPalettes) Set(k syntaxhighlight.Kind, v termbox.Attribute) {
	c[k] = v
}

func (c ColorPalettes) Get(k syntaxhighlight.Kind) termbox.Attribute {
	// ignore whitespace kind
	if k == syntaxhighlight.Whitespace {
		return termbox.ColorRed
	}

	v, ok := c[k]
	if !ok {
		panic(fmt.Sprintf("Unknown syntax highlight kind %d\n", k))
	}

	return v
}
