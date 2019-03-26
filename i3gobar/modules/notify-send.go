package modules

import (
	"fmt"
	"os/exec"
)

func NotifySend(format string, a ...interface{}) {
	Notify(format, a...).AppName("i3gobar").Send()
}

type Dunst map[string]string

func Notify(format string, a ...interface{}) Dunst {
	d := Dunst{}
	d["message"] = fmt.Sprintf(format, a...)
	return d
}

// AppName -a, --appname=NAME Name of your application
func (d Dunst) AppName(name string) Dunst {
	d["appname"] = name
	return d
}

// -u, --urgency=URG           The urgency of this notification
// -h, --hints=HINT            User specified hints
// -A, --action=ACTION         Actions the user can invoke
func (d Dunst) Action(action, label string) Dunst {
	d["action"] = action + "," + label
	return d
}

// -t, --timeout=TIMEOUT       The time until the notification expires
// -i, --icon=ICON             An Icon that should be displayed with the notification
// -I, --raw_icon=RAW_ICON     Path to the icon to be sent as raw image data
// -c, --capabilities          Print the server capabilities and exit
// -s, --serverinfo            Print server information and exit
// -p, --printid               Print id, which can be used to update/replace this notification
// -r, --replace=ID            Set id of this notification.
func (d Dunst) Replace(id int) Dunst {
	d["replace"] = fmt.Sprint(id)
	return d
}

// -C, --close=ID              Set id of this notification.
// -b, --block                 Block until notification is closed and print close reason

func (d Dunst) Send() {
	args := []string{}
	for flag, value := range d {
		if flag == "message" {
			continue
		}
		args = append(args, "--"+flag, value)
	}
	args = append(args, d["message"])
	exec.Command("dunstify", args...).Run()
}
