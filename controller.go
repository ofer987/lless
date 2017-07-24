package main

import (
	"github.com/nsf/termbox-go"
)

func handle(fpath string, colorPalettes ColorPalettes) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	s, err := NewStream(fpath, colorPalettes)
	if err != nil {
		panic(err)
	}
	s.Render()

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
				case ev.Key == termbox.KeyArrowDown || ev.Ch == 'j':
					s.moveDown()
					s.Render()
				case ev.Key == termbox.KeyArrowUp || ev.Ch == 'k':
					s.moveUp()
					s.Render()
				case ev.Key == termbox.KeyArrowLeft || ev.Ch == 'h':
					s.moveLeft()
					s.Render()
				case ev.Key == termbox.KeyArrowRight || ev.Ch == 'l':
					s.moveRight()
					s.Render()
				case ev.Key == termbox.KeyEsc || ev.Ch == 'q' || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
					return
				}
			}
		}
	}
}
