package main

import (
	"time"

	"github.com/zwzn/i3/i3gobar/bar"
	"github.com/zwzn/i3/i3gobar/modules"
)

func main() {

	bar.Run(
		bar.Schedule(modules.Weather, time.Hour),
		bar.Schedule(modules.ZFS, time.Hour),
		bar.Schedule(modules.Pamac, time.Hour).OnClick(func(click bar.Click) {
			modules.Notify(modules.Shell("pamac checkupdates")()).AppName("pamac").Send()
		}),
		bar.Schedule(modules.Shell("task list"), time.Second*10),
		bar.Schedule(modules.Shell("$HOME/bin/battery"), time.Second*10),
		bar.Schedule(modules.Prepend("MST ", modules.TimeIn("America/Phoenix")), time.Second),
		bar.Schedule(modules.DateTime, time.Second),
	)
}
