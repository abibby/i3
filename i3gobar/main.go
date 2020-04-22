package main

import (
	"os/exec"
	"time"

	"github.com/abibby/i3/i3gobar/bar"
	"github.com/abibby/i3/i3gobar/modules"
)

func main() {

	bar.Run(
		// modules.Timer(),
		bar.Schedule(modules.Weather, time.Hour).OnClick(func(click bar.Click) {
			if click.Button != bar.MouseLeft {
				return
			}
			modules.Shell("xdg-open 'https://weather.gc.ca/forecast/hourly/on-5_metric_e.html'")()
		}),
		modules.Music(),
		bar.Schedule(modules.ZFS, time.Hour),
		bar.Schedule(modules.Mail, time.Minute),
		bar.Schedule(modules.Pamac, time.Hour).OnClick(func(click bar.Click) {
			b, _ := exec.Command("pamac", "checkupdates").Output()
			modules.Notify(string(b)).AppName("pamac").Send()
		}),
		bar.Schedule(modules.Shell("$HOME/bin/battery"), time.Second*10),
		bar.Schedule(modules.DateTime, time.Second),
	)
}
