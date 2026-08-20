# XML Converter

![CI](https://github.com/Qyroxen/XML-Converter/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/XML-Converter/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/XML-Converter?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/XML-Converter)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/XML-Converter)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/XML-Converter?style=social)](https://github.com/Qyroxen/XML-Converter/stargazers)

## What is it?

XML Converter is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/XML-Converter.git
cd XML-Converter
go build -o xmlconverter .

# Run
./xmlconverter --help
```

## CLI Usage

```bash
# Basic usage
./xmlconverter

# With flags
./xmlconverter --verbose --output json

# Get help
./xmlconverter --help
```

## Examples

```bash
# Example 1
./xmlconverter example1

# Example 2
./xmlconverter example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o xmlconverter .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/XML-Converter/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/XML-Converter?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/XML-Converter/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/XML-Converter?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/XML-Converter/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/XML-Converter" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/XML-Converter/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/XML-Converter" alt="Pull Requests">
  </a>
</p>
