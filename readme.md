## KNFGen

`KNFGen` is utility for generating go const code for [KNF](https://godoc.org/pkg.re/essentialkaos/ek.v1/knf) configs.

* [Installation](#installation)
* [Usage](#usage)
* [Build Status](#build-status)
* [Contributing](#contributing)
* [License](#license)

#### Installation

To build the KNFGen from scratch, make sure you have a working Go 1.5+ workspace ([instructions](https://golang.org/doc/install)), then:

```
go get github.com/essentialkaos/knfgen
```

If you want update KNFGen to latest stable release, do:

```
go get -u github.com/essentialkaos/knfgen
```

#### Usage

```
Usage: knfgen <options> config-file
    
Options:
    
  --separators, -s     Add new lines between sections
  --no-color, -nc      Disable colors in output
  --help, -h           Show this help message
  --version, -v        Show version

```

#### Build Status

| Repository | Status |
|------------|--------|
| Stable | [![Build Status](https://travis-ci.org/essentialkaos/knfgen.svg?branch=master)](https://travis-ci.org/essentialkaos/knfgen) |
| Unstable | [![Build Status](https://travis-ci.org/essentialkaos/knfgen.svg?branch=develop)](https://travis-ci.org/essentialkaos/knfgen) |

#### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

#### License

[EKOL](https://essentialkaos.com/ekol)
