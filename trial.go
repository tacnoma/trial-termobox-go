package main

import (
  "time"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
  time.Sleep(1 * time.Second)
	defer termbox.Close()
}
