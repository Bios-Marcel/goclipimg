package goclipimg

import (
	"C"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	kernel32 = windows.NewLazyDLL("kernel32.dll")
	user32   = windows.NewLazyDLL("user32.dll")

	// kernel32.dll
	kernel32GlobalLock   = kernel32.NewProc("GlobalLock")
	kernel32GlobalSize   = kernel32.NewProc("GlobalSize")
	kernel32GlobalUnlock = kernel32.NewProc("GlobalUnlock")

	// user32.dll
	user32OpenClipboard    = user32.NewProc("OpenClipboard")
	user32GetDesktopWindow = user32.NewProc("GetClipboardOwner")
	user32GetClipboardData = user32.NewProc("GetClipboardData")
	user32CloseClipboard   = user32.NewProc("CloseClipboard")
)

// kernel32.dll

func globalLock(hMem uintptr) (uintptr, error) {
	result, _, findError := kernel32GlobalLock.Call(hMem)
	if result == 0 {
		return 0, findError
	}

	return result, nil
}

func globalSize(hMem uintptr) (int, error) {
	result, _, findError := kernel32GlobalSize.Call(hMem)
	if result == 0 {
		return 0, findError
	}
	return int(result), nil
}

func globalUnlock(hMem uintptr) error {
	success, _, findError := kernel32GlobalUnlock.Call(hMem)
	if success == 0 {
		return findError
	}

	return nil
}

// user32.dll

func getDesktopWindow() (uintptr, error) {
	result, _, findError := user32GetDesktopWindow.Call()
	if result == 0 {
		return 0, findError
	}

	return result, nil
}

func openClipboard(hWndNewOwner uintptr) error {
	success, _, findError := user32OpenClipboard.Call(hWndNewOwner)
	if success == 0 {
		return findError
	}

	return nil
}

func getClipboardData(uFormat uint) (uintptr, error) {
	result, _, findError := user32GetClipboardData.Call(uintptr(uFormat))
	if result == 0 {
		return 0, findError
	}

	return result, nil
}

func closeClipboard() error {
	success, _, findError := user32CloseClipboard.Call()
	if success == 0 {
		return findError
	}

	return nil
}

func getImageFromClipboard() ([]byte, error) {

	openClipboardError := openClipboard(uintptr(0))
	if openClipboardError != nil {
		return nil, openClipboardError
	}

	defer closeClipboard()

	clipboardObjectHandle, getClipboardDataError := getClipboardData(17)
	if getClipboardDataError != nil {
		return nil, getClipboardDataError
	}

	clipboardObject, globalLockError := globalLock(clipboardObjectHandle)
	if globalLockError != nil {
		return nil, globalLockError
	}

	clipboardObjectSize, globalSizeError := globalSize(clipboardObjectHandle)
	if globalSizeError != nil {
		return nil, globalSizeError
	}

	result := C.GoBytes(unsafe.Pointer(clipboardObject), C.int(clipboardObjectSize))

	globalUnlock(clipboardObjectHandle)

	return result, nil
}
