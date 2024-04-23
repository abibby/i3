package modules

import (
	"fmt"

	"github.com/abibby/i3/i3gobar/icon"
)

func ZFS() string {
	status := Shell(`ssh frank "zpool list -o name,cap,health" | grep volume1 | awk '{ print $2 " " $3 }'`)()

	return fmt.Sprintf("%s %s", icon.HDD, status)
}
