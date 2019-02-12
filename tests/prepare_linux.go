package tests

import (
	"os"
	"os/exec"
)

const imageName = "image.png"

func fillClipboard() error {
	_, statError := os.Stat(imageName)
	if statError != nil {
		return statError
	}

	return exec.Command("xclip", "-sel", "clipboard", imageName).Run()
}
