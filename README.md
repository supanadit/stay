# _GOSTAY_

This is an alternative of `go get`, the story why i build this software because `go get` not showing specific progress while
downloading package

## Compatibility

- Windows ( NO )
- Linux ( YES )
- Mac OS ( It should be work too, Not Tested )

## Install
### Method 1 ( Still using `go get` )
1. Using go get -v - u github.com/supanadit/gostay
2. Go to $GOPATH/src/github.com/supanadit/gostay
3. `go install`

Make sure your $GOPATH/bin include to Environment such as `.bash_profile`

### Method 2
1. Download [release](https://github.com/supanadit/gostay/releases) of this Project
2. Copy `gostay` to `/usr/bin` or `$GOPATH/bin`

## How it Works

After you installing this app, just use command
```shell script
gostay -u <your_package>
```
![Go Stay Example](https://i.ibb.co/tHRqBNS/gostay.png)

## Feature
- Download Golang package [OK]

## TODO
- Create web interface for golang package manager Web GUI

## Thanks To

[https://github.com/akamensky/argparse](https://github.com/akamensky/argparse)

[https://gopkg.in/src-d/go-git.v4](https://gopkg.in/src-d/go-git.v4)

## License

Copyright 2019 Supan Adit Pratama

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
