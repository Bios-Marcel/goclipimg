# goclipimg

![tests](https://github.com/github/docs/actions/workflows/test.yml/badge.svg)

This is just a tiny library that helps you getting an image from your
clipboard into your application.

## Requirements

### Linux

If you are running X11 you will need to have `xclip` installed.

On Wayland you will need `wl-clipboard`.

### Windows

On Windows, this simply invokes a PowerShell script. Sadly this causes the
library to open a PowerShell window, unless you are already running inside of
a terminal.

## Example

```go
func main() {
    data, err := goclipimg.GetImageFromClipboard()
    if err == nil {
        doSomethingWithYourPNG(data)
    }
}
```