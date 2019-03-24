package bar

import (
	"time"
)

type Click struct {
	Button    int      `json:"button"`
	Modifiers []string `json:"modifiers"`
	X         int      `json:"x"`
	Y         int      `json:"y"`
	RelativeX int      `json:"relative_x"`
	RelativeY int      `json:"relative_y"`
	Width     int      `json:"width"`
	Height    int      `json:"height"`
}

type Block struct {
	onClick func(button Click)
	text    chan string
}

func (b *Block) OnClick(cb func(button Click)) *Block {
	b.onClick = cb
	return b
}

func Schedule(cb func() string, every time.Duration) *Block {
	cs := make(chan string)
	go func() {
		cs <- cb()
		for range time.Tick(every) {
			cs <- cb()
		}
	}()
	return &Block{
		text: cs,
	}
}