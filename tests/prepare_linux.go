package tests

import "os/exec"

func fillClipboard() error {
	return exec.Command("xclip", "-sel", "clipboard", "image.png").Run()
}
