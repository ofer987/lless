package main

import (
	"github.com/nsf/termbox-go"
)

func displayLoop(fpath string, colorPalettes ColorPalettes) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)

	s, err := NewStream(fpath, colorPalettes)
	if err != nil {
		panic(err)
	}
	defer s.CloseStream()
	t := s.Read()
	t.Render()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Ch == 'g':
					t.goToFirstLine()
					t.Render()
				case ev.Ch == 'G':
					t.goToLastLine()
					t.Render()
				case ev.Ch == 'd' || ev.Key == termbox.KeyCtrlD:
					t.jumpDown()
					t.Render()
				case ev.Ch == 'u' || ev.Key == termbox.KeyCtrlU:
					t.jumpUp()
					t.Render()
				case ev.Key == termbox.KeyArrowDown || ev.Ch == 'j':
					t.moveDown()
					t.Render()
				case ev.Key == termbox.KeyArrowUp || ev.Ch == 'k':
					t.moveUp()
					t.Render()
				case ev.Key == termbox.KeyArrowLeft || ev.Ch == 'h':
					t.moveLeft()
					t.Render()
				case ev.Key == termbox.KeyArrowRight || ev.Ch == 'l':
					t.moveRight()
					t.Render()
				case ev.Key == termbox.KeyEsc || ev.Ch == 'q' || ev.Key == termbox.KeyCtrlC:
					return
				}
			}
		}
	}
}
