// +build !linux

package goclipimg

// GetImageFromClipboard always returns ErrImagePasteUnsupported, since the
// compilation target currently doesn't support pasting images.
func GetImageFromClipboard() ([]byte, error) {
	return nil, ErrImagePasteUnsupported
}
