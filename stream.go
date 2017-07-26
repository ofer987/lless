package main

import (
	"fmt"
	"io"
	"os"

	"github.com/nsf/termbox-go"
	"github.com/sourcegraph/syntaxhighlight"
)

type Stream struct {
	reader        *os.File
	colorPalettes ColorPalettes
	text          Text
}

func NewStdinStream(colorPalettes ColorPalettes) (*Stream, error) {
	file := os.Stdin
	if file == nil {
		return nil, fmt.Errorf("cannot read from standard in")
	}

	s := &Stream{reader: file, colorPalettes: colorPalettes}
	return s, nil
}

func NewFileStream(fpath string, colorPalettes ColorPalettes) (*Stream, error) {
	file, err := os.Open(fpath)

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

	s := &Stream{
		reader:        file,
		colorPalettes: colorPalettes,
	}
	return s, nil
}

func (s *Stream) Read() Text {
	s.text = Text{lines: make([]Line, 1)}

	syntaxhighlight.Print(
		syntaxhighlight.NewScannerReader(s.reader),
		nil,
		s,
	)

	return s.text
}

func (s *Stream) Close() error {
	return s.reader.Close()
}

func (s *Stream) Print(_ io.Writer, kind syntaxhighlight.Kind, tokText string) error {
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
			s.text.lines[len(s.text.lines)-1] = append(*lastLine, Char{c, s.get(kind)})
		}
	}

	return nil
}

func (s *Stream) get(k syntaxhighlight.Kind) termbox.Attribute {
	// ignore whitespace kind
	if k == syntaxhighlight.Whitespace {
		return termbox.ColorRed
	}

	v, ok := s.colorPalettes[k]
	if !ok {
		panic(fmt.Sprintf("Unknown syntax highlight kind %d\n", k))
	}

	return v
}
