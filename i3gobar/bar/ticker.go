package bar

import (
	"strings"
	"time"
)

func Ticker(length int, in *Block) *Block {
	cs := make(chan string)
	go func() {
		out := ""
		i := 0
		for {
			select {
			case text := <-in.text:
				out = text
				i = 0
			case <-time.Tick(time.Millisecond * 250):
				if len(out) < length {
					cs <- out
				} else {
					cs <- strings.Join(strings.Split(out+"    "+out, "")[i:i+length], "")
					i++
					i = i % (len(out) + 4)
				}
			}
		}
	}()
	return &Block{
		text: cs,
	}
}
