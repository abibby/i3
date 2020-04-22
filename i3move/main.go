package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/davecgh/go-spew/spew"
	"github.com/abibby/i3"
)

var spaces []i3.Workspace

func main() {
	var up bool
	var down bool
	var left bool
	var right bool
	var all bool

	flag.BoolVar(&up, "up", false, "move to the previous workspace on this monitor")
	flag.BoolVar(&down, "down", false, "move to the next workspace on this monitor")
	flag.BoolVar(&left, "left", false, "move to the workspace on the monitor to the left")
	flag.BoolVar(&right, "right", false, "move to the workspace on the monitor to the right")

	flag.BoolVar(&all, "all", false, "move all workspaces")

	flag.Parse()

	var err error
	spaces, err = i3.GetWorkspaces()
	check(err)

	if up {
		if num, ok := spaceUp(); ok {
			i3.Command("workspace %d", num)
		}
	}

	if down {
		if num, ok := spaceDown(); ok {
			i3.Command("workspace %d", num)
		}
	}
	if left {
		if num, ok := spaceLeft(); ok {
			i3.Command("workspace %d", num)
		}
	}
	if right {
		if num, ok := spaceRight(); ok {
			i3.Command("workspace %d", num)
		}
	}
}

func visible() []i3.Workspace {
	vis := []i3.Workspace{}
	for _, space := range spaces {
		if space.Visible {
			vis = append(vis, space)
		}
	}
	nm := len(vis)
	sort.Slice(vis, func(i, j int) bool {
		return (vis[i].Num-1)%nm < (vis[j].Num-1)%nm
	})
	return vis
}

func numMonitors() int {
	return len(visible())
}
func monitor() int {
	vis := visible()
	cs := currentSpace()
	for i, space := range vis {
		if space.Num == cs.Num {
			return i
		}
	}
	panic("no visible monitors")
}

func currentSpace() i3.Workspace {
	for _, space := range spaces {
		if space.Focused {
			return space
		}
	}
	panic("no focused space")
}

func spaceUp() (int, bool) {
	num := currentSpace().Num - numMonitors()
	return num, validSpace(num)
}

func spaceDown() (int, bool) {
	num := currentSpace().Num + numMonitors()
	return num, validSpace(num)
}

func changeMonitor(i int) (int, bool) {
	cm := monitor() + i
	vis := visible()
	if cm < 0 {
		return 0, false
	}
	if cm >= len(vis) {
		return 0, false
	}
	spew.Dump(cm, vis)
	return vis[cm].Num, true
}

func spaceLeft() (int, bool) {
	return changeMonitor(-1)
}

func spaceRight() (int, bool) {
	return changeMonitor(1)
}

func validSpace(num int) bool {
	if num < 0 {
		return false
	}
	if num > 9 {
		return false
	}
	return true
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
