package goclipimg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func GetImageFromClipboard() ([]byte, error) {
	tempFile, tempFileError := ioutil.TempFile("", "clipimg.jpg")
	if tempFileError != nil {
		return nil, tempFileError
	}

	imagePath := filepath.Join(os.TempDir(), tempFile.Name())
	pasteError := exec.Command("pngpaste", imagePath).Run()
	if pasteError != nil {
		return nil, errors.New(fmt.Sprintf("error pasting image into temporary file: %s", pasteError.Error()))
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
