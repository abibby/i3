package main

import (
	"time"

	"github.com/zwzn/i3/i3gobar/modules"
)

func main() {
	run(
		schedule(modules.Weather, time.Hour),
		schedule(modules.ZFS, time.Hour),
		schedule(modules.Pamac, time.Hour),
		schedule(modules.Shell("$HOME/bin/i3blocks-contrib/battery2/battery2"), time.Second*10),
		schedule(modules.DateTime, time.Second),
	)
}
