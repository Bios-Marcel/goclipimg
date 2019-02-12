package tests

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/Bios-Marcel/goclipimg"
)

func TestGetImageFromClipboard(t *testing.T) {
	testData, readError := ioutil.ReadFile("image.png")
	if readError != nil {
		t.Fatalf("Error reading test data: %s", readError.Error())
	}

	if len(testData) != 40230 {
		t.Errorf("Incorrect test data, length should have been 40230, but was %d", len(testData))
	}

	runError := fillClipboard()
	if runError != nil {
		t.Errorf("Error getting testdata into clipboard: %s", runError.Error())
	}

	data, readError := goclipimg.GetImageFromClipboard()
	if readError != nil {
		t.Fatalf("Error reading image from clipboard: %s", readError.Error())
	}

	if bytes.Compare(testData[:], data[:]) != 0 {
		t.Fatalf("Data is incorrect. LenTestData: %d | LenClipData: %d", len(testData), len(data))
	}
}
