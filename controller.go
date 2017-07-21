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

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	s, err := NewStream(fpath, colorPalettes)
	if err != nil {
		panic(err)
	}

	for {
		s.Render()
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowDown || ev.Ch == 'j':
					s.moveDown()
				case ev.Key == termbox.KeyArrowUp || ev.Ch == 'k':
					s.moveUp()
				case ev.Key == termbox.KeyArrowLeft || ev.Ch == 'h':
					s.moveLeft()
				case ev.Key == termbox.KeyArrowRight || ev.Ch == 'l':
					s.moveRight()
				case ev.Key == termbox.KeyEsc || ev.Ch == 'q' || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
					return
				}
			}
		}
	}
}
