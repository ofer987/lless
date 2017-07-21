package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/nsf/termbox-go"
	"github.com/sourcegraph/syntaxhighlight"
)

var x, y int

var stdout io.Writer

type Stream struct {
	filePath      string
	colorPalettes ColorPalettes
}

func NewStream(fpath string, colorPalettes ColorPalettes) (*Stream, error) {
	s := new(Stream)

	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	stdout = colorable.NewColorableStdout()

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if fi.Mode().IsDir() {
		return nil, fmt.Errorf("%s is a directory", file.Name())
	}

	s.colorPalettes = colorPalettes
	s.filePath = fpath

	x = 0
	y = 0

	return s, nil
}

func (s *Stream) Render() {
	file, err := os.Open(s.filePath)
	if err != nil {
		panic(err)
	}

	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Mode().IsDir() {
		panic(fmt.Errorf("%s is a directory", file.Name()))
	}

	termbox.Clear(termbox.ColorWhite, termbox.ColorWhite)

	// x = 0
	// y = 0

	// originalX := x

	// for _, c := range s.contents {
	// 	termbox.SetCell(x, y, c, termbox.ColorBlack, termbox.ColorWhite)
	// 	x++
	// }
	// x = originalX
	termbox.Flush()
	syntaxhighlight.Print(
		syntaxhighlight.NewScannerReader(file),
		stdout,
		s,
	)
}

func (s *Stream) moveDown() {
	y--
}

func (s *Stream) moveUp() {
	y++
}

func (s *Stream) moveLeft() {
	x--
}

func (s *Stream) moveRight() {
	x++
}

func (s *Stream) Print(w io.Writer, kind syntaxhighlight.Kind, tokText string) error {
	color := s.colorPalettes.Get(kind)

	// termbox.SetCell(x, y, 'c', termbox.ColorRed, termbox.ColorWhite)
	originalX := x
	originalY := y
	for _, c := range tokText {
		if c == '\n' {
			x = 0
			y++
		} else {
			termbox.SetCell(x, y, c, color, termbox.ColorWhite)
			x++
		}
	}
	x = originalX
	y = originalY

	termbox.Flush()
	return nil
}
