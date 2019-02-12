package goclipimg

import "errors"

var (
	// ErrNoImageInClipboard means that no data was returned.
	ErrNoImageInClipboard = errors.New("the clipboard doesn't contain an image")
)
