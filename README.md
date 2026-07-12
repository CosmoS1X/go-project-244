# Differ

[![Go](https://github.com/CosmoS1X/differ/actions/workflows/go.yml/badge.svg)](https://github.com/CosmoS1X/differ/actions/workflows/go.yml)

## Overview

A small command-line utility for comparing two configuration files and showing the differences between them.

## Features

- Supports JSON and YAML files
- Supports both absolute and relative file paths
- Outputs differences in three formats: stylish, plain, and JSON
- Can be used as a library from Go code

## Installation

You can install the application using Go:

```bash
go install github.com/CosmoS1X/differ/cmd/differ@latest
```

You can also download a precompiled binary from the [releases page]().

## Usage

```bash
differ <file1> <file2>
```

Optional flags:

By default, the output format is stylish.

The format flag supports both short and long forms:

```bash
differ --format stylish <file1> <file2>
differ -f plain <file1> <file2>
differ -f json <file1> <file2>
```

## Library usage

The package can also be used as a Go library. Import the package and call the diff function:

```go
import "github.com/CosmoS1X/differ"

result, err := differ.Gen("file1.json", "file2.json", "stylish")
```
