# goclipimg

This is just a tiny library that helps you getting an image from your clipboard into your application.

## CI

**CI can not be run on a headless linux server, as the clipboard is part of
the XServer, which isn't running on a headless vm**

I'll try adding CI for MacOS and Windows, in case it is possible.

## Example

```go
func main() {
    data, readError := goclipimg.GetImageFromClipboard()
    if readError == nil {
        doSomethingWithYourPNG(data)
    }
}
```