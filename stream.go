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
	text          Text
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

	x = 0
	y = 0

	return s, nil
}

func (s *Stream) Read() *Text {
	termbox.Clear(termbox.ColorWhite, termbox.ColorWhite)
	termbox.Flush()

	s.text = Text{lines: make([]Line, 1)}

	syntaxhighlight.Print(
		syntaxhighlight.NewScannerReader(s.reader),
		nil,
		s,
	)

	return &s.text
}

func (s *Stream) CloseStream() {
}

func (s *Stream) Print(w io.Writer, kind syntaxhighlight.Kind, tokText string) error {
	for _, c := range tokText {
		if c == '\n' {
			newLine := Line{}
			s.text.lines = append(s.text.lines, newLine)
		} else if c == '\t' {
			i := 0
			lastLine := s.text.lastLine()
			for i < 4 {
				s.text.lines[len(s.text.lines)-1] = append(*lastLine, Char{' ', termbox.ColorDefault})
				i++
			}
		} else {
			lastLine := s.text.lastLine()
			s.text.lines[len(s.text.lines)-1] = append(*lastLine, Char{c, s.colorPalettes.Get(kind)})
		}
	}

	return nil
}
