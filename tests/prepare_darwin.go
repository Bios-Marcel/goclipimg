package tests

import "os/exec"

const imageName = "image.jpg"
const osaScript = "set the clipboard to (read (POSIX file \"" + imageName + "\") as JPEG picture)"

func fillClipboard() error {
	return exec.Command("osascript", "-e", osaScript).Run()
}
