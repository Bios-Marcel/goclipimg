package tests

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/Bios-Marcel/goclipimg"
)

func TestGetImageFromClipboard(t *testing.T) {
	data, readError := goclipimg.GetImageFromClipboard()
	if readError != goclipimg.ErrNoImageInClipboard {
		t.Fatalf("no data was put into clipboard, error was %s; data was %x", readError, data)
	}

	testData, readError := ioutil.ReadFile(imageName)
	if readError != nil {
		t.Fatalf("Error reading test data: %s", readError.Error())
	}

	if len(testData) != 40230 {
		t.Fatalf("Incorrect test data, length should have been 40230, but was %d", len(testData))
	}

	fillError := fillClipboard()
	if fillError != nil {
		t.Fatalf("Error getting testdata into clipboard: %s", fillError.Error())
	}

	data, readError = goclipimg.GetImageFromClipboard()
	if readError != nil {
		t.Fatalf("Error reading image from clipboard: %s", readError.Error())
	}

	if bytes.Compare(testData[:], data[:]) != 0 {
		t.Fatalf("Data is incorrect. LenTestData: %d | LenClipData: %d", len(testData), len(data))
	}
}
