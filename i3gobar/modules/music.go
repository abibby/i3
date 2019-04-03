package modules

import (
	"fmt"
	"log"

	"github.com/zwzn/gpmdp"
	"github.com/zwzn/i3/i3gobar/bar"
	"github.com/zwzn/i3/i3gobar/icon"
)

func Music() *bar.Block {
	cs := make(chan string)

	go func() {
		g, err := gpmdp.Connect()
		if err != nil {
			log.Fatal(err)
		}
		playing := false
		track := &gpmdp.Track{}
		for {
			select {
			case track = <-g.Track():
			case playing = <-g.Playing():
			}

			ico := icon.PauseCircle
			if playing {
				ico = icon.PlayCircle
			}

			cs <- fmt.Sprintf("%s %s - %s\n", ico, track.Title, track.Artist)
		}

	}()
	return bar.NewBlock(cs)
}
