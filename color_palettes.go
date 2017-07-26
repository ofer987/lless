package main

import (
	"github.com/nsf/termbox-go"
	"github.com/sourcegraph/syntaxhighlight"
)

var (
	LightColorPalettes = ColorPalettes{
		syntaxhighlight.String:        termbox.Attribute(58),
		syntaxhighlight.Keyword:       termbox.Attribute(21),
		syntaxhighlight.Comment:       termbox.Attribute(7),
		syntaxhighlight.Type:          termbox.Attribute(36),
		syntaxhighlight.Literal:       termbox.Attribute(36),
		syntaxhighlight.Punctuation:   termbox.ColorRed,
		syntaxhighlight.Plaintext:     termbox.Attribute(21),
		syntaxhighlight.Tag:           termbox.ColorBlue,
		syntaxhighlight.HTMLTag:       termbox.Attribute(82),
		syntaxhighlight.HTMLAttrName:  termbox.ColorBlue,
		syntaxhighlight.HTMLAttrValue: termbox.ColorGreen,
		syntaxhighlight.Decimal:       termbox.ColorBlue,
	}

	DarkColorPalettes = ColorPalettes{
		syntaxhighlight.String:        termbox.Attribute(58),
		syntaxhighlight.Keyword:       termbox.ColorBlue,
		syntaxhighlight.Comment:       termbox.Attribute(243),
		syntaxhighlight.Type:          termbox.Attribute(75),
		syntaxhighlight.Literal:       termbox.Attribute(75),
		syntaxhighlight.Punctuation:   termbox.ColorRed,
		syntaxhighlight.Plaintext:     termbox.ColorBlue,
		syntaxhighlight.Tag:           termbox.ColorBlue,
		syntaxhighlight.HTMLTag:       termbox.Attribute(82),
		syntaxhighlight.HTMLAttrName:  termbox.ColorBlue,
		syntaxhighlight.HTMLAttrValue: termbox.ColorGreen,
		syntaxhighlight.Decimal:       termbox.ColorBlue,
	}
)

type ColorPalettes map[syntaxhighlight.Kind]termbox.Attribute
