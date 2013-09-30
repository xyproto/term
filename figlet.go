package textgui

// Figlet wrapper

import (
	"os/exec"
)

const (
	figletcmd = "/usr/bin/figlet"
)

// Uses figlet to generate ascii art text
func Figlet(msg string) (string, error) {
	cmd := exec.Command(figletcmd, msg)
	b, err := cmd.Output()
	if err != nil {
		return msg, err
	}
	return string(b), nil
}
