package main

import (
	"time"

	"github.com/abibby/i3/i3gobar/bar"
	"github.com/abibby/i3/i3gobar/modules"
)

func main() {
	// spew.Dump(modules.GitHubNotifications())
	// modules.JiraNotifications()
	// os.Exit(3)
	bar.Run(
		// modules.Timer(),
		bar.Schedule(modules.Weather, time.Hour).OnClick(func(click bar.Click) {
			if click.Button != bar.MouseLeft {
				return
			}
			modules.Shell("xdg-open 'https://weather.gc.ca/forecast/hourly/on-5_metric_e.html'")()
		}),
		// modules.Music(),
		bar.Schedule(modules.GitHubNotifications, time.Minute*1).OnClick(func(click bar.Click) {
			if click.Button != bar.MouseLeft {
				return
			}

			modules.Shell("xdg-open 'https://github.com/notifications'")()
		}),
		// bar.Schedule(modules.JiraNotifications, time.Minute*1),
		bar.Schedule(modules.ZFS, time.Hour),
		// bar.Schedule(modules.Mail, time.Minute),
		// bar.Schedule(modules.Pamac, time.Hour).OnClick(func(click bar.Click) {
		// 	b, _ := exec.Command("pamac", "checkupdates").Output()
		// 	modules.Notify(string(b)).AppName("pamac").Send()
		// }),
		// bar.Schedule(modules.Shell("$HOME/bin/battery"), time.Second*10),
		bar.Schedule(modules.DateTime, time.Second),
	)
}
