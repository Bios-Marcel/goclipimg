# goclipimg

| OS | Build |
| - | - |
| linux | [![builds.sr.ht status](https://builds.sr.ht/~biosmarcel/goclipimg/arch.yml.svg)](https://builds.sr.ht/~biosmarcel/goclipimg/arch.yml?) |
| darwin | [![Build Status](https://travis-ci.org/Bios-Marcel/goclipimg.svg?branch=master)](https://travis-ci.org/Bios-Marcel/goclipimg) |
| windows | **TODO** |

This is just a tiny library that helps you getting an image from your clipboard into your application.

## Requirements

### Requirements - Linux

Currently only the xserver is supported, therefore `xserver` and `xclip` need to be installed.

## Example

```go
func main() {
    data, readError := goclipimg.GetImageFromClipboard()
    if readError == nil {
        doSomethingWithYourPNG(data)
    }
}
```