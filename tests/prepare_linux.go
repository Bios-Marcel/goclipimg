package tests

import (
	"os"
	"os/exec"
)

func fillClipboard() error {
	_, statError := os.Stat("image.png")
	if statError != nil {
		return statError
	}

	return exec.Command("xclip", "-sel", "clipboard", "image.png").Run()
}
