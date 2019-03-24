package main

import (
	"time"

	"github.com/zwzn/i3/i3gobar/bar"
	"github.com/zwzn/i3/i3gobar/modules"
)

func main() {
	// go func() {
	// 	f, err := os.OpenFile("/home/adam/go/src/github.com/zwzn/i3/i3gobar/bar/output.txt", os.O_CREATE|os.O_RDWR, 0644)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	io.Copy(f, os.Stdin)
	// }()

	bar.Run(
		bar.Schedule(modules.Weather, time.Hour),
		bar.Schedule(modules.ZFS, time.Hour),
		bar.Schedule(modules.Pamac, time.Hour).OnClick(func(click bar.Click) {
			modules.NotifySend("%#v", click)
		}),
		bar.Schedule(modules.Shell("$HOME/bin/i3blocks-contrib/battery2/battery2"), time.Second*10),
		bar.Schedule(modules.DateTime, time.Second),
	)
}
