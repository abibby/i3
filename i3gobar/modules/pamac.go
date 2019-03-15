package modules

import (
	"fmt"
)

func Pamac() string {
	count := Shell("pamac checkupdates -q | wc -l")()

	if count == "0" {
		return ""
	}

	return fmt.Sprintf("<span font='Font Awesome'></span> %s", count)
}
