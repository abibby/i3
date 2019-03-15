package modules

import (
	"os/exec"
)

func Shell(cmd string) func() string {
	return func() string {

		b, _ := exec.Command("bash", "-c", cmd).Output()

		return string(b)
	}
}
