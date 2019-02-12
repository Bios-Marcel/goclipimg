package tests

import "os/exec"

const imageName = "image.png"
const osaScript = "set the clipboard to (read (POSIX file '" + imageName + "') as «class PNGf»)"

func fillClipboard() error {
	return exec.Command("osascript", "-e", osaScript).Run()
}
