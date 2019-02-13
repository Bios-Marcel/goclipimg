package goclipimg

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

const imageToFile = "tell application \"System Events\" to write (the clipboard as «class PNGf») to \"%s\""
const containsImageLineOne = "if ((clipboard info) as string) does not contain \"«class PNGf»\" then"
const containsImageLineTwo = "error number 61"
const containsImageLineThree = "end if"

// ErrImagePasteUnsupported means that pngpaste can't be found or isn't installed.
var ErrImagePasteUnsupported = errors.New("pngpaste can't be found on this system")

func getImageFromClipboard() ([]byte, error) {
	containsImageCommand := exec.Command("osascript", "-s", "o", "-e", containsImageLineOne, "-e", containsImageLineTwo, "-e", containsImageLineThree)
	errorPipeContainsImage, err := containsImageCommand.StdoutPipe()
	if err != nil {
		return nil, err
	}

	waitForContainsCheckError := containsImageCommand.Run()
	if waitForContainsCheckError != nil {
		if exitError, ok := waitForContainsCheckError.(*exec.ExitError); ok {
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				if status.ExitStatus() == 61 {
					return nil, ErrNoImageInClipboard
				}
			}
		}

		errScanner := bufio.NewScanner(errorPipeContainsImage)
		errOut := ""
		for errScanner.Scan() {
			errOut = errOut + errScanner.Text()
		}

		return nil, errors.New(fmt.Sprintf("Error checking for png: %s. - %s", waitForContainsCheckError.Error(), errOut))
	}

	tempFile, tempFileError := ioutil.TempFile("", "clipimg")
	if tempFileError != nil {
		return nil, tempFileError
	}

	imagePath := tempFile.Name()
	defer os.Remove(imagePath)

	command := exec.Command("osascript", "-e", fmt.Sprintf(imageToFile, imagePath))

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
