package tests

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
)

const imageName = "image.png"
const osaScript = "set the clipboard to (read (POSIX file '" + imageName + "') as PNG)"

func fillClipboard() error {
	command := exec.Command("osascript", "-e", osaScript)

	errorPipe, err := command.StderrPipe()
	if err != nil {
		return err
	}

	outputPipe, err := command.StdoutPipe()
	if err != nil {
		return err
	}

	commandError := command.Start()
	if commandError != nil {
		return commandError
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
		return errors.New(fmt.Sprintf("%s | %s", errOut, out))
	}

	return nil
}
