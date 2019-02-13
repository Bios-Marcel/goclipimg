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
const containsImage = "if ((clipboard info) as string) does not contain \"«class PNGf»\" then error number 61 end if"

// ErrImagePasteUnsupported means that pngpaste can't be found or isn't installed.
var ErrImagePasteUnsupported = errors.New("pngpaste can't be found on this system")

func getImageFromClipboard() ([]byte, error) {
	containsImageCommand := exec.Command("osascript", "-e", containsImage)

	containsCheckError := containsImageCommand.Start()
	if containsCheckError != nil {
		return nil, containsCheckError
	}

	waitForContainsCheckError := containsImageCommand.Wait()
	if waitForContainsCheckError != nil {
		if exitError, ok := waitForContainsCheckError.(*exec.ExitError); ok {
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				if status == 61 {
					return nil, ErrNoImageInClipboard
				}
			}
		} else {
			return nil, waitForContainsCheckError
		}
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
