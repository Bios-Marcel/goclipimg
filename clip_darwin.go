package goclipimg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const imageToFile = "tell application \"System Events\" to write (the clipboard as «class PNGf») to \"%s\""

// ErrImagePasteUnsupported means that pngpaste can't be found or isn't installed.
var ErrImagePasteUnsupported = errors.New("pngpaste can't be found on this system")

func getImageFromClipboard() ([]byte, error) {
	tempFile, tempFileError := ioutil.TempFile("", "clipimg")
	if tempFileError != nil {
		return nil, tempFileError
	}

	imagePath := tempFile.Name()
	defer os.Remove(imagePath)

	command := exec.Command("osascript", "-e", fmt.Sprintf(imageToFile, imagePath))

	commandError := command.Run()
	if commandError != nil {
		return nil, ErrNoImageInClipboard
	}

	data, readError := ioutil.ReadFile(imagePath)
	if readError != nil {
		return nil, readError
	}

	if len(data) == 0 {
		return nil, ErrNoImageInClipboard
	}

	return data, nil
}
