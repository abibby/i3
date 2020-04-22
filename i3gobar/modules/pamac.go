package modules

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/abibby/i3/i3gobar/icon"
)

func Pamac() string {
	count, _ := strconv.Atoi(strings.TrimSpace(Shell("pamac checkupdates -q | wc -l")()))

	if count == 0 {
		return ""
	}

	return fmt.Sprintf("%s %d", icon.Box, count)
}
