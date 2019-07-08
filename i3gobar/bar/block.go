package bar

import (
	"time"
)

type MouseButton uint8

const (
	MouseLeft       = 1
	MouseMiddle     = 2
	MouseRight      = 3
	MouseScrollUp   = 4
	MouseScrollDown = 5
	MouseBack       = 8
	MouseForward    = 9
)

type Click struct {
	Name      string      `json:"name"`
	Instance  string      `json:"instance"`
	Button    MouseButton `json:"button"`
	Modifiers []string    `json:"modifiers"`
	X         int         `json:"x"`
	Y         int         `json:"y"`
	RelativeX int         `json:"relative_x"`
	RelativeY int         `json:"relative_y"`
	Width     int         `json:"width"`
	Height    int         `json:"height"`
}

type Block struct {
	onClick func(button Click)
	text    chan string
}

func NewBlock(cs chan string) *Block {
	return &Block{
		text: cs,
	}
}

func (b *Block) OnClick(cb func(Click)) *Block {
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
