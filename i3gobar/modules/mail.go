package modules

import (
	"fmt"
	"os/exec"

	"github.com/abibby/i3/i3gobar/icon"
	_ "github.com/mattn/go-sqlite3"
)

var lastList = ""

func Mail() string {
	count := Shell("mailq count")()
	if count == "0" {
		return ""
	}
	b, err := exec.Command("mailq", "list").Output()
	if err != nil {
		return err.Error()
	}
	list := string(b)
	if list != lastList {
		Notify(list).AppName("mail").Send()
	}

	return fmt.Sprintf("%s %s", icon.Envelope, count)
}
