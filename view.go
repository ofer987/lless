package main

import (
	// "fmt"
	// "math"
	// "strings"
	// "bytes"
	"io"
	"io/ioutil"

	"github.com/nsf/termbox-go"
	"github.com/sourcegraph/syntaxhighlight"
)

// Colours
const backgroundColour = termbox.ColorWhite
const textColour = termbox.ColorBlue

// type ColourReader struct {
// 	ColorPalettes ColorPalettes
// 	Reader        io.Reader
// }

type TbPrinter struct {
	ColorPalettes ColorPalettes
	X             int
	Y             int
}

func render(r io.Reader, w io.Writer, cp ColorPalettes) error {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	termbox.Clear(backgroundColour, backgroundColour)
	x = 0
	y = 0

	syntaxhighlight.Print(
		syntaxhighlight.NewScannerReader(r),
		w,
		TbPrinter{cp, 0, 0},
	)
	// error checking

	// to flush or not to flush?
	// execute in ensure block?

	// which error to return?
	return nil
}

func render2(r io.Reader, w io.Writer, cp ColorPalettes) error {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	// termbox.Clear(backgroundColour, backgroundColour)

	byteArray, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	// n := bytes.IndexByte(byteArray, 0)
	s := string(byteArray[:])
	// runes := []rune(s)

	x := 0
	y := 0

	// var s string
	// s = "foobar"
	for _, r := range s {
		if r == '\n' {
			x = 0
			y++
		} else {
			termbox.SetCell(x, y, r, termbox.ColorBlack, termbox.ColorWhite)
			x++
		}

		// if x%60 == 0 {
		// 	x = 0
		// 	y++
		// }
	}
	termbox.Flush()

	if false {
		syntaxhighlight.Print(
			syntaxhighlight.NewScannerReader(r),
			w,
			TbPrinter{cp, 0, 0},
		)
	}
	defer termbox.Close()
	// error checking

	// to flush or not to flush?
	// execute in ensure block?

	// which error to return?
	return nil
}

func (p TbPrinter) Print(w io.Writer, kind syntaxhighlight.Kind, tokText string) error {
	color := p.ColorPalettes.Get(kind)

	for _, c := range tokText {
		if c == '\n' {
			x = 0
			y++
		} else {
			termbox.SetCell(x, y, c, color, termbox.ColorWhite)
			x++
		}
	}

	termbox.Flush()
	return nil
}
