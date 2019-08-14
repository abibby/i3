package modules

import (
	"time"

	"github.com/zwzn/i3/i3gobar/bar"
)

func Timer() *bar.Block {
	cs := make(chan string)
	reset := make(chan struct{})

	go func() {
		startTime := time.Now()
		cs <- time.Since(startTime).String()
		for {
			select {
			case <-time.Tick(time.Second):
			case <-reset:
				startTime = time.Now()
			}
			cs <- time.Since(startTime).Round(time.Second).String()
		}
	}()

	return bar.NewBlock(cs).OnClick(func(click bar.Click) {
		if click.Button == bar.MouseLeft {
			reset <- struct{}{}
		}
	})
}
