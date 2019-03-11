package i3

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func Msg(args ...string) (json.RawMessage, error) {
	return exec.Command("i3-msg", args...).Output()
}

func Command(cmd string, a ...interface{}) error {
	data, err := Msg(fmt.Sprintf(cmd, a...))
	if err != nil {
		return err
	}
	resp := []struct {
		Success bool `json:"success"`
	}{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}
	if len(resp) == 0 {
		return fmt.Errorf("no responce")
	}
	if !resp[0].Success {
		return fmt.Errorf("command failed")
	}
	return nil
}

type Rect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Workspace struct {
	Num     int    `json:"num"`
	Name    string `json:"name"`
	Visible bool   `json:"visible"`
	Focused bool   `json:"focused"`
	Rect    Rect   `json:"rect"`
	Output  string `json:"output"`
	Urgent  bool   `json:"urgent"`
}

func GetWorkspaces() ([]Workspace, error) {
	b, err := Msg("-t", "get_workspaces")
	if err != nil {
		return nil, err
	}
	spaces := []Workspace{}
	err = json.Unmarshal(b, &spaces)
	if err != nil {
		return nil, err
	}

	return spaces, nil
}
