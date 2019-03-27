package modules

import (
	"time"
)

func DateTime() string {
	return time.Now().Format("Mon 2 3:04PM")
}
func TimeIn(location string) func() string {
	return func() string {
		loc, err := time.LoadLocation(location)
		if err != nil {
			panic(err)
		}
		return time.Now().In(loc).Format("3:04PM")
	}
}
