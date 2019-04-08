package modules

import (
	"fmt"
	"html"

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
			return
		}
		ico := icon.PauseCircle
		title := ""
		for {
			select {
			case err := <-g.Error:
				cs <- err.Error()
				g.Close()
				return
			case ev := <-g.Event:
				if track, ok := ev.Track(); ok {
					title = html.EscapeString(fmt.Sprintf("%s - %s\n", track.Title, track.Artist))

				}
				if playing, ok := ev.Playing(); ok {
					if playing {
						ico = icon.Pause
					} else {
						ico = icon.Play
					}
				}

				cs <- fmt.Sprintf("%s %s %s %s\n", icon.StepBackward, ico, icon.StepForward, title)
			}
		}

	}()
	return bar.NewBlock(cs).OnClick(func(click bar.Click) {
		if click.RelativeX < 16 {
			Shell("playerctl previous")()
		} else if click.RelativeX < 42 {
			Shell("playerctl play-pause")()
		} else if click.RelativeX < 62 {
			Shell("playerctl next")()
		}
		// Notify(spew.Sdump(click)).Replace(10341).Send()
	})
}
