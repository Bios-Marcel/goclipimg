// +build linux

package goclipimg

import (
	"bytes"
	"io"
	"os/exec"
)

func isCommandAvailable(name string) bool {
	_, fileError := exec.LookPath(name)
	return fileError == nil
}

func GetImageFromClipboard() ([]byte, error) {
	if !isCommandAvailable("xclip") {
		return nil, ErrImagePasteUnsupported
	}

	xclip := exec.Command("xclip", "-sel", "clipboard", "-t", "image/png", "-o")
	output, outputError := xclip.StdoutPipe()
	if outputError != nil {
		return nil, outputError
	}

	startError := xclip.Start()
	if startError != nil {
		return nil, startError
	}

	data := make([]byte, 0)
	buffer := bytes.NewBuffer(data)
	io.Copy(buffer, output)

	image := buffer.Bytes()
	if len(image) == 0 {
		return nil, ErrNoImageInClipboard
	}

	return image, nil
}
