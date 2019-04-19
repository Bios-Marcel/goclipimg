package tests

import (
	"os/exec"
	"testing"

	"github.com/Bios-Marcel/goclipimg"
)

func TestGetImageFromClipboard(t *testing.T) {
	fillError := fillClipboard()
	if fillError != nil {
		t.Fatalf("Error getting testdata into clipboard: %s", fillError.Error())
	}

	defer exec.Command("powershell", "-Command", "Set-Clipboard").Run()

	data, readError := goclipimg.GetImageFromClipboard()
	if readError != nil {
		t.Fatalf("Error reading image from clipboard: %s", readError.Error())
	}

	// at least as big as a png header, not really a good test, but hey.
	if len(data) <= 8 {
		t.Fatal("Data is incorrect, length was 0")
	}
}

func TestGetImageFromClipboardNoImagePresent(t *testing.T) {
	_, imageError := goclipimg.GetImageFromClipboard()
	if imageError != goclipimg.ErrNoImageInClipboard {
		t.Errorf("Expected error was 'ErrNoImageInClipboard', but got %s", imageError.Error())
	}
}
