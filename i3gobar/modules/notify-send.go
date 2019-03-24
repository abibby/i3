package modules

import "fmt"

func NotifySend(format string, a ...interface{}) {
	Shell(fmt.Sprintf("notify-send '%s'", fmt.Sprintf(format, a...)))()
}
