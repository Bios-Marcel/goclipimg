package goclipimg

import (
	"bytes"
	"errors"
)

var (
	pngHeader = []byte{
		/* Has the high bit set to detect transmission systems that do not
		support 8 bit data and to reduce the chance that a text file i
		mistakenly interpreted as a PNG, or vice versa. */
		0x89,
		/* In ASCII, the letters PNG, allowing a person to identify the
		   format easily if it is viewed in a text editor. */
		0x50, 0x4E, 0x47,
		/* A DOS-style line ending (CRLF) to detect DOS-Unix line ending
		   conversion of the data.  */
		0x0D, 0x0A,
		/* A byte that stops display of the file under DOS when the command
		   type has been used—the end-of-file character.  */
		0x1A,
		/* A Unix-style line ending (LF) to detect Unix-DOS line ending
		   conversion.  */
		0x0A}
	// ErrNoImageInClipboard means that no data was returned.
	ErrNoImageInClipboard = errors.New("the clipboard doesn't contain an image")
)

// GetImageFromClipboard returns either a byte array containing PNG data or an
// error that indicates that no png could be retrieved.
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
