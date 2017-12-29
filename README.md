# Scanner

[![Build Status](https://travis-ci.org/cpliakas/fileinfo.svg?branch=master)](https://travis-ci.org/cpliakas/fileinfo)
[![GoDoc](https://godoc.org/github.com/cpliakas/fileinfo?status.svg)](https://godoc.org/github.com/cpliakas/fileinfo)
[![Go Report Card](https://goreportcard.com/badge/github.com/cpliakas/fileinfo)](https://goreportcard.com/report/github.com/cpliakas/fileinfo)

A go package that extracts and stores basic metadata about a file.

## Installation

Assuming a [correctly configured](https://golang.org/doc/install#testing) Go
toolchain:

```shell
go get github.com/cpliakas/fileinfo
```

## Usage

The code below writes a file's MD5 sum to STDOUT.

```go
package main

import (
	"fmt"

	"github.com/cpliakas/fileinfo"
)

func main() {
	i, err := fileinfo.New("/path/to/file")
	if err != nil {
		panic(err)
    }

    hash, _ := i.Hash()
	fmt.Println(hash)
}
```