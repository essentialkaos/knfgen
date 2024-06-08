<p align="center"><a href="#readme"><img src="https://gh.kaos.st/knfgen.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/w/knfgen/ci"><img src="https://kaos.sh/w/knfgen/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/knfgen/codeql"><img src="https://kaos.sh/w/knfgen/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="https://kaos.sh/r/knfgen"><img src="https://kaos.sh/r/knfgen.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/b/knfgen"><img src="https://kaos.sh/b/3ae560e1-1fef-4ca7-b46a-17558e105963.svg" alt="Codebeat badge" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`knfgen` is utility for generating Go const code for [KNF](https://kaos.sh/knf-spec) configuration files.

### Installation

#### From source

To build the `knfgen` from scratch, make sure you have a working Go 1.21+ workspace ([instructions](https://go.dev/doc/install)), then:

```
go install github.com/essentialkaos/knfgen@latest
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux and macOS from [EK Apps Repository](https://apps.kaos.st/knfgen/latest).

To install the latest prebuilt version of `knfgen`, do:

```bash
bash <(curl -fsSL https://apps.kaos.st/get) knfgen
```

### Usage

<img src=".github/images/usage.svg" />

### Build Status

| Branch | Status |
|------------|--------|
| `master` | [![CI](https://kaos.sh/w/knfgen/ci.svg?branch=master)](https://kaos.sh/w/knfgen/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/knfgen/ci.svg?branch=develop)](https://kaos.sh/w/knfgen/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
