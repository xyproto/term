package textgui

// Dialog wrapper

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"bytes"
)

const (
	dialogcmd = "/usr/bin/dialog"
)

type Dialog struct {
	width  int
	height int
}

func NewDialog(width, height int) *Dialog {
	return &Dialog{width, height}
}

// Uses dialog to display a message box
func (d *Dialog) MsgBox(msg string) {
	cmd := exec.Command(dialogcmd, "--msgbox", msg, strconv.Itoa(d.height), strconv.Itoa(d.width))
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (d *Dialog) Menu(text string, menuheight int, menuItemPairs map[string]string) string {
	args := []string{"--menu", text, strconv.Itoa(d.height), strconv.Itoa(d.width), strconv.Itoa(menuheight)}
	for key, value := range menuItemPairs {
		args = append(args, key)
		args = append(args, value)
	}
	cmd := exec.Command(dialogcmd, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	var out bytes.Buffer
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return out.String()
}
