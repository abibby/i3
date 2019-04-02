package modules

import (
	"os/exec"
	"strings"
)

func Shell(cmd string) func() string {
	return func() string {

		b, _ := exec.Command("bash", "-c", cmd).Output()

		lines := strings.Split(strings.TrimSpace(string(b)), "\n")
		return lines[len(lines)-1]
	}
}
