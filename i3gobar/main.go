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
		bar.Schedule(modules.Pamac, time.Hour),
		bar.Schedule(modules.Shell("$HOME/bin/i3blocks-contrib/battery2/battery2"), time.Second*10),
		bar.Schedule(modules.DateTime, time.Second),
	)
}
