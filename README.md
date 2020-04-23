# STAY
Bulletproof package manager fo `Go`.

## How it works

- Installing and download one package
```shell script
stay -i github.com/gin-gonic/gin
```

- Installing and download the package also find related package to it. flag `-a` meaning `auto`
```shell script
stay -g github.com/gin-gonic/gin
```

or

```shell script
stay --get-related github.com/gin-gonic/gin
```

- Installing package more than one, with `requirements.txt`
```shell script
stay -f requirements.txt
```

- Uninstall or Remove package
```shell script
stay -r github.com/supanadit/devops-factory
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

Make sure your $GOPATH/bin include to Environment such as `.bash_profile`

### Method 2
1. Download [release](https://github.com/supanadit/stay/releases) of this Project
2. Copy `stay` to `/usr/bin` or `$GOPATH/bin`

## Feature
- Download Golang package
- Install and Download Package like Python by `stay -f requirements.txt`
- Get Package and find related package by `stay -a github.com/gin-gonic/gin` or `stay --get-related github.com/gin-gonic/gin`
- Remove Package or Uninstall Package by `stay -r github.com/supanadit/devops-factory`

## TODO
- Create web interface for golang package manager Web GUI
- Support downloading package from golang.org
- Automatically Check Needed Dependency
- Update Package
- Show all installed package
- Support downloading private package use SSH Auth
- Go Modules Support

## Thanks To

- [https://github.com/akamensky/argparse](https://github.com/akamensky/argparse)
- [https://gopkg.in/src-d/go-git.v4](https://gopkg.in/src-d/go-git.v4)

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
