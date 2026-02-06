# Clipboard for Go

[![GoDoc](https://godoc.org/github.com/zxyqq/clipboard?status.svg)](http://godoc.org/github.com/zxyqq/clipboard)

Cross-platform clipboard package that works on Windows, WSL, macOS and Linux.

```bash
go get github.com/zxyqq/clipboard
```

## Usage


```go
clipboard.WriteAll("Hello World!")
```

```go
text, err := clipboard.ReadAll()
```

## License

[MIT](LICENSE)
