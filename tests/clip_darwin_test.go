package tests

import (
	"testing"

	"github.com/Bios-Marcel/goclipimg"
)

func TestGetImageFromClipboard(t *testing.T) {
	fillError := fillClipboard()
	if fillError != nil {
		t.Fatalf("Error getting testdata into clipboard: %s", fillError.Error())
	}

	data, readError := goclipimg.GetImageFromClipboard()
	if readError != nil {
		t.Fatalf("Error reading image from clipboard: %s", readError.Error())
	}

	// at least as big as a png header, not really a good test, but hey.
	if len(data) <= 8 {
		t.Fatal("Data is incorrect, length was 0")
	}
}
