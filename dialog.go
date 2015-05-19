package term

// Dialog wrapper

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
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

	// Sort the keys
	keys := make([]string, len(menuItemPairs))
	i := 0
	for key, _ := range menuItemPairs {
		keys[i] = key
		i++
	}
	sort.Strings(keys)

	// Append the menu entries in sorted order
	for _, key := range keys {
		args = append(args, key)
		args = append(args, menuItemPairs[key])
	}

	// Run the dialog tool
	cmd := exec.Command(dialogcmd, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	var out bytes.Buffer
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Return the output from the dialog tool
	return out.String()
}
