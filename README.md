# Stay

[![Go Report Card](https://goreportcard.com/badge/github.com/supanadit/stay)](https://goreportcard.com/report/github.com/supanadit/stay)

Bulletproof package manager for Go.

## How it works

```shell script
stay -i github.com/gin-gonic/gin
```

## Compatibility

- Windows ( Coming soon )
- Linux
- Mac OS

## Install

### Method 1 ( Still using `go get` )
1. Using `go get -u -v github.com/supanadit/stay`
2. Go to $GOPATH/src/github.com/supanadit/stay
3. `go install`

Make sure your $GOPATH/bin added into environment variable such as `.bash_profile`, `.profile`, or `.bashrc`

### Method 2
1. Download the [release](https://github.com/supanadit/stay/releases) version
2. Copy `stay` to `/usr/bin` or `$GOPATH/bin`

## Thanks To
- [https://github.com/akamensky/argparse](https://github.com/akamensky/argparse)
- [https://gopkg.in/src-d/go-git.v4](https://gopkg.in/src-d/go-git.v4)
