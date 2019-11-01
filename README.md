![Logo](http://supanadit.com/wp-content/uploads/2019/11/Gostay-Logo.png)

# _GOSTAY_
This is an alternative of `go get`, the story why i build this software because `go get` not showing specific progress while
downloading package

## How it works

![Go Stay Example](https://i.ibb.co/tHRqBNS/gostay.png)

- Installing and download one package
```shell script
gostay -u github.com/gin-gonic/gin
```

- Installing and download the package also find related package to it. flag `-a` meaning `auto`
```shell script
gostay -a github.com/gin-gonic/gin
```

or

```shell script
gostay --get-related github.com/gin-gonic/gin
```

- Installing package more than one, with `requirements.txt`
```shell script
gostay -f requirements.txt
```

- Uninstall or Remove package
```shell script
gostay -r github.com/supanadit/devops-factory
```

## Compatibility

- Windows
- Linux
- Mac OS

## Install
### Method 1 ( Still using `go get` )
1. Using `go get -u -v github.com/supanadit/gostay`
2. Go to $GOPATH/src/github.com/supanadit/gostay
3. `go install`

Make sure your $GOPATH/bin include to Environment such as `.bash_profile`

### Method 2
1. Download [release](https://github.com/supanadit/gostay/releases) of this Project
2. Copy `gostay` to `/usr/bin` or `$GOPATH/bin`

## Feature
- Download Golang package [OK]
- Install and Download Package like Python using `pip install -r requirements.txt` [OK]
- Get Package and find related package by `gostay -a github.com/gin-gonic/gin` or `gostay --get-related github.com/gin-gonic/gin` [OK]
- Remove Package or Uninstall Package by `gostay -r github.com/supanadit/devops-factory` [OK]

## TODO
- Create web interface for golang package manager Web GUI
- Support downloading package from golang.org
- Automatically Check Needed Dependency
- Update Package
- Show all installed package
- Support downloading private package use SSH Auth

## Changelog
### Version 1.3
- Legacy Support and could find related package
- Remove Package or Uninstall Package
### Version 1.2
- Support All Operating System
### Version 1.1
- Can automatically installing and download using `requirement.txt` with flag `-f` it similar like `requirement.txt` in Python 
### Version 1.0
- Can download package with flag `-u`

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
