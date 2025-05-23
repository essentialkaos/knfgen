<p align="center"><a href="#readme"><img src=".github/images/card.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/w/knfgen/ci"><img src="https://kaos.sh/w/knfgen/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/knfgen/codeql"><img src="https://kaos.sh/w/knfgen/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src=".github/images/license.svg"/></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#ci-status">CI Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`knfgen` is utility for generating Go const code for [KNF](https://kaos.sh/knf-spec) configuration files.

### Installation

#### From source

To build the `knfgen` from scratch, make sure you have a working [Go 1.23+](https://github.com/essentialkaos/.github/blob/master/GO-VERSION-SUPPORT.md) workspace ([instructions](https://go.dev/doc/install)), then:

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

<img src=".github/images/usage.svg"/>

### CI Status

| Branch | Status |
|------------|--------|
| `master` | [![CI](https://kaos.sh/w/knfgen/ci.svg?branch=master)](https://kaos.sh/w/knfgen/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/knfgen/ci.svg?branch=develop)](https://kaos.sh/w/knfgen/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/.github/blob/master/CONTRIBUTING.md).

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://kaos.dev"><img src="https://raw.githubusercontent.com/essentialkaos/.github/refs/heads/master/images/ekgh.svg"/></a></p>
