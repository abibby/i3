package modules

import (
	"fmt"
	"html"
	"log"

	"github.com/abibby/gpmdp"
	"github.com/abibby/i3/i3gobar/bar"
	"github.com/abibby/i3/i3gobar/icon"
)

func Music() *bar.Block {
	cs := make(chan string)

	go func() {
		g, err := gpmdp.Connect()
		if err != nil {
			log.Println(err)
			cs <- ""
			return
		}
		ico := ""
		title := ""
		for {
			select {
			case err := <-g.Error:
				log.Println(err)
				cs <- ""
				g.Close()
				return
			case ev := <-g.Event:
				if track, ok := ev.Track(); ok {
					title = html.EscapeString(fmt.Sprintf(
						"%s - %s\n",
						truncate(track.Title, 25),
						truncate(track.Artist, 20),
					))

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
