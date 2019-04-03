package modules

import (
	"fmt"

	"github.com/zwzn/gpmdp"
	"github.com/zwzn/i3/i3gobar/bar"
	"github.com/zwzn/i3/i3gobar/icon"
)

func Music() *bar.Block {
	cs := make(chan string)

	go func() {
		g, err := gpmdp.Connect()
		if err != nil {
			cs <- err.Error()
		}
		ico := icon.PauseCircle
		title := ""
		for {
			select {
			case err := <-g.Error:
				cs <- err.Error()

			case ev := <-g.Event:
				if track, ok := ev.Track(); ok {
					title = fmt.Sprintf("%s - %s\n", track.Title, track.Artist)
				}
				if playing, ok := ev.Playing(); ok {
					if playing {
						ico = icon.PlayCircle
					} else {
						ico = icon.PauseCircle
					}
				}

				cs <- fmt.Sprintf("%s %s\n", ico, title)
			}
		}

	}()
	return bar.NewBlock(cs)
}
