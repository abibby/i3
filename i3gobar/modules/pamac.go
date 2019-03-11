package modules

import (
	"fmt"
	"os/exec"
	"strings"
)

func Pamac() string {
	b, _ := exec.Command("pamac", "checkupdates", "-q").Output()

	count := strings.Count(string(b), "\n")

	return fmt.Sprintf("pkg %d", count)
}
