package tests

import "os/exec"

const osaScript = "set the clipboard to (read (POSIX file \"image.png\") as PNG picture)"

func fillClipboard() error {
	return exec.Command("osascript", "-e", osaScript).Run()
}
