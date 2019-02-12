# goclipimg

[![builds.sr.ht status](https://builds.sr.ht/~biosmarcel/goclipimg/arch.yml.svg)](https://builds.sr.ht/~biosmarcel/goclipimg/arch.yml?)

This is just a tiny library that helps you getting an image from your clipboard into your application.

## Example

```go
func main() {
    data, readError := goclipimg.GetImageFromClipboard()
    if readError == nil {
        doSomethingWithYourPNG(data)
    }
}
```