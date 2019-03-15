package modules

import (
	"os/exec"
	"strings"
)

func Shell(cmd string) func() string {
	return func() string {

		b, _ := exec.Command("bash", "-c", cmd).Output()

		parts := strings.Split(strings.TrimRight(string(b), "\n"), "\n")
		return parts[len(parts)-1]
	}
}
