Simple command to generate a permalink from the source code shown by godoc on golang.org to the repository viewer at code.google.com/p/go

## Install
```
># go install github.com/ghthor/godocpermalink
```

## Usage

```
># godocpermalink http://golang.org/src/pkg/net/url/url.go?s=5522:5827#L220
https://code.google.com/p/go/source/browse/src/pkg/net/url/url.go?name=go1.3.1#220
>#
```
