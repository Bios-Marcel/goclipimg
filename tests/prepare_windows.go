package tests

import (
	"fmt"
	"os"
	"os/exec"
)

const imageName = "image.png"

func fillClipboard() error {
	_, statError := os.Stat(imageName)
	if statError != nil {
		return statError
	}

	return exec.Command("powershell", "-Command", fmt.Sprintf(`Add-Type -Assembly System.Windows.Forms
	Add-Type -Assembly System.Drawing	
	$img = [Drawing.Image]::FromFile('%s')
	[Windows.Forms.Clipboard]::SetImage($img)
	$img.Dispose()`, imageName)).Run()
}
