package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/nsf/termbox-go"
	"github.com/sourcegraph/syntaxhighlight"
)

var x, y int

type Stream struct {
	reader        *bytes.Reader
	colorPalettes ColorPalettes
	lenLines      int
}

func NewStream(fpath string, colorPalettes ColorPalettes) (*Stream, error) {
	s := new(Stream)

	file, err := os.Open(fpath)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if fi.Mode().IsDir() {
		return nil, fmt.Errorf("%s is a directory", file.Name())
	}

	fileContents, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	s.reader = bytes.NewReader(fileContents)
	s.colorPalettes = colorPalettes
	s.lenLines = lenLines(string(fileContents))

	x = 0
	y = 0

	return s, nil
}

func (s *Stream) Render() {
	originalX := x
	originalY := y

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
		syntaxhighlight.NewScannerReader(s.reader),
		nil,
		s,
	)
	if _, err := s.reader.Seek(0, 0); err != nil {
		panic(err)
	}

	x = originalX
	y = originalY
}

func (s *Stream) CloseStream() {
}

func (s *Stream) jumpUp() {
	remainingLines := 0 - y

	if remainingLines >= 10 {
		y += 10
	} else if remainingLines < 10 && remainingLines > 0 {
		y += remainingLines
	}
}

func (s *Stream) jumpDown() {
	remainingLines := s.lenLines + y

	if remainingLines >= 10 {
		y -= 10
	} else if remainingLines < 10 && remainingLines > 0 {
		y -= remainingLines
	}
}

func (s *Stream) moveDown() {
	if s.lenLines+y > 0 {
		y--
	}
}

func (s *Stream) moveUp() {
	if y < 0 {
		y++
	}
}

func (s *Stream) moveLeft() {
	if x > 0 {
		x--
	}
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

func lenLines(s string) int {
	result := 0
	for _, c := range s {
		if c == '\n' {
			result++
		}
	}

	return result
}
