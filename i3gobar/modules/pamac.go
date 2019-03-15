package modules

import (
	"fmt"

	"github.com/zwzn/i3/i3gobar/icon"
)

func Pamac() string {
	count := Shell("pamac checkupdates -q | wc -l")()

	if count == "0" {
		return ""
	}

	return fmt.Sprintf("%s %s", icon.Box, count)
}
