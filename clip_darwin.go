package goclipimg

import (
	"bufio"
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

	command := exec.Command("pngpaste", imagePath)

	errorPipe, err := command.StderrPipe()
	if err != nil {
		return nil, err
	}

	outputPipe, err := command.StdoutPipe()
	if err != nil {
		return nil, err
	}

	commandError := command.Start()
	if commandError != nil {
		return nil, commandError
	}

	errScanner := bufio.NewScanner(errorPipe)
	errOut := ""
	for errScanner.Scan() {
		errOut = errOut + errScanner.Text()
	}

	outScanner := bufio.NewScanner(outputPipe)
	out := ""
	for outScanner.Scan() {
		out = out + outScanner.Text()
	}

	if errOut != "" || out != "" {
		return nil, errors.New(fmt.Sprintf("%s | %s", errOut, out))
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
