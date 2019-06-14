<p align="center"><a href="#readme"><img src="https://gh.kaos.st/knfgen.svg"/></a></p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<p align="center">
  <a href="https://travis-ci.org/essentialkaos/knfgen"><img src="https://travis-ci.org/essentialkaos/knfgen.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/knfgen"><img src="https://goreportcard.com/badge/github.com/essentialkaos/knfgen"></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-knfgen-master"><img alt="codebeat badge" src="https://codebeat.co/badges/3ae560e1-1fef-4ca7-b46a-17558e105963" /></a>
  <a href="https://essentialkaos.com/ekol"><img src="https://gh.kaos.st/ekol.svg"></a>
</p>

`KNFGen` is utility for generating Go const code for [KNF](https://godoc.org/pkg.re/essentialkaos/ek.v10/knf) configs.

### Installation

#### From source

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)):

```
git config --global http.https://pkg.re.followRedirects true
```

To build the KNFGen from scratch, make sure you have a working Go 1.10+ workspace ([instructions](https://golang.org/doc/install)), then:

```
go get github.com/essentialkaos/knfgen
```

If you want to update KNFGen to latest stable release, do:

```
go get -u github.com/essentialkaos/knfgen
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux and OS X from [EK Apps Repository](https://apps.kaos.st/knfgen/latest).

### Usage

```
Usage: knfgen {options} config-file
    
Options:
    
  --separators, -s     Add new lines between sections
  --no-color, -nc      Disable colors in output
  --help, -h           Show this help message
  --version, -v        Show version

```

### Build Status

| Branch | Status |
|------------|--------|
| `master` | [![Build Status](https://travis-ci.org/essentialkaos/knfgen.svg?branch=master)](https://travis-ci.org/essentialkaos/knfgen) |
| `develop` | [![Build Status](https://travis-ci.org/essentialkaos/knfgen.svg?branch=develop)](https://travis-ci.org/essentialkaos/knfgen) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[EKOL](https://essentialkaos.com/ekol)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>