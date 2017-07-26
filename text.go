package main

import (
	"github.com/nsf/termbox-go"
)

type Text struct {
	lines []Line
	x     int
	y     int
}

type Line []Char

type Char struct {
	r     rune
	color termbox.Attribute
}

func (t Text) lastLine() *Line {
	return &t.lines[len(t.lines)-1]
}

func (t Text) Render() {
	termbox.Clear(termbox.ColorWhite, termbox.ColorWhite)

	for y, line := range t.lines[t.y:] {
		for x, char := range line[t.x:] {
			termbox.SetCell(x, y, char.r, char.color, termbox.ColorWhite)
		}
	}

	termbox.Flush()
}

func (t *Text) jumpUp() {
	remainingLines := t.y

	if remainingLines >= 10 {
		t.y -= 10
	} else if remainingLines > 0 && remainingLines < 10 {
		t.y -= remainingLines
	}
}

func (t *Text) jumpDown() {
	remainingLines := len(t.lines) - t.y

	if remainingLines >= 10 {
		t.y += 10
	} else if remainingLines > 5 && remainingLines < 10 {
		t.y += remainingLines - 5
	}
}

func (t *Text) moveDown() {
	if t.y < len(t.lines)-2 {
		t.y++
	}
}

func (t *Text) moveUp() {
	if t.y > 0 {
		t.y--
	}
}

func (t *Text) moveLeft() {
}

func (t *Text) moveRight() {
}
