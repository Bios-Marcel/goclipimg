//go:build windows

package goclipimg

import (
	"fmt"
	"os"
	"os/exec"
)

func getImageFromClipboard() ([]byte, error) {
	tempFile, tempFileError := os.CreateTemp("", "clipimg")
	if tempFileError != nil {
		return nil, tempFileError
	}

	imagePath := tempFile.Name()

	tempFile.Close()
	defer os.Remove(imagePath)

	err := exec.Command("powershell", "-Command", fmt.Sprintf(`Add-Type -Assembly PresentationCore
	$img = [Windows.Clipboard]::GetImage()
	if (!($img -eq $null)) {
		$file = '%s'
		$stream = [IO.File]::Open($file, 'OpenOrCreate')
		$encoder = New-Object Windows.Media.Imaging.PngBitmapEncoder
		$encoder.Frames.Add([Windows.Media.Imaging.BitmapFrame]::Create($img))
		$encoder.Save($stream)
		$stream.Dispose()
	}`, imagePath)).Run()
	if err != nil {
		return nil, ErrNoImageInClipboard
	}

	data, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, ErrNoImageInClipboard
	}

	return data, nil
}
