package tests

import (
	"bufio"
	"os"
	"os/exec"
)

const imageName = "image.png"

func fillClipboard() error {
	_, statError := os.Stat(imageName)
	if statError != nil {
		return statError
	}

	sessionType := os.Getenv("XDG_SESSION_TYPE")
	if sessionType == "wayland" {
		copy := exec.Command("wl-copy")
		file, err := os.Open(imageName)
		if err != nil {
			return err
		}
		copy.Stdin = bufio.NewReader(file)
		return copy.Run()
	}

	return exec.Command("xclip", "-sel", "clipboard", imageName).Run()
}
