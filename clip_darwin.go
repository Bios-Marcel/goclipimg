package goclipimg

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func GetImageFromClipboard() ([]byte, error) {
	tempFile, tempFileError := ioutil.TempFile("", "clipimg.png")
	if tempFileError != nil {
		return nil, tempFileError
	}

	imagePath := filepath.Join(os.TempDir(), tempFile.Name())
	pasteError := exec.Command("pngpaste", imagePage).Run()
	if pasteError != nil {
		return errors.New("error pasting image into temporary file: %s", pasteError.Error())
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
