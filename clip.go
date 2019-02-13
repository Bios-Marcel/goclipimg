package goclipimg

import (
	"bytes"
	"errors"
)

var (
	pngHeader = []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	// ErrNoImageInClipboard means that no data was returned.
	ErrNoImageInClipboard = errors.New("the clipboard doesn't contain an image")
)

func GetImageFromClipboard() ([]byte, error) {
	data, err := getImageFromClipboard()
	if err != nil {
		return nil, err
	}

	if len(data) < 8 {
		return nil, ErrNoImageInClipboard
	}

	if bytes.Compare(data[:8], pngHeader[:]) != 0 {
		return nil, ErrNoImageInClipboard
	}

	return data, nil
}
