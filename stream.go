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
	defer file.Close()

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
	originalX := x
	originalY := y

	file, err := os.Open(s.filePath)
	defer file.Close()

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

	x = originalX
	y = originalY
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

	for _, c := range tokText {
		if c == '\n' {
			x = 0
			y++
		} else if c == '\t' {
			x += 4
		} else {

			termbox.SetCell(x, y, c, color, termbox.ColorWhite)
			x++
		}
	}

	termbox.Flush()
	return nil
}
