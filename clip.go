package goclipimg

import "errors"

var (
	// ErrImagePasteUnsupported means that xclip can't be found or isn't installed.
	ErrImagePasteUnsupported = errors.New("xclip is nto available on this system")
	// ErrNoImageInClipboard means that no data was returned.
	ErrNoImageInClipboard = errors.New("the clipboard doesn't contain an image")
)
