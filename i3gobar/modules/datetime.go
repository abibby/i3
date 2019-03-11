package modules

import (
	"time"
)

func DateTime() string {
	return time.Now().Format("Mon 2 3:04PM")
}
