package modules

import "fmt"

func ZFS() string {
	status := Shell(`ssh frank-wan "zpool list -o name,cap,health" | grep volume1`)()

	return fmt.Sprintf("<span font='Font Awesome'></span> %s", status)
}
