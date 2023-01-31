# The simple logger for golang  

[![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Go Report Card][report-img]][report]

### The simple extension of the standard "log" library to minimize effort, code and ease of use

![lg][example-img]

The logger uses the standard go "log" package. Allows you to install and use expanded information (prefix, date and time, file and line) without cluttering up with unnecessary code. It also colors text on Linux and OS X for better reading of logs.  

## Install

```bash
go get -u github.com/bearatol/lg
```

## Exampels

```go
package main

import (
	"github.com/bearatol/lg"
)

func main() {
	testArray := [...]int{1, 2, 3, 4, 5}

	// Info
	lg.Info("test text", testArray)
	lg.Infoln("test text", testArray)
	lg.Infof("some values: %s, %v", "test text", testArray)

	// Debug
	lg.Debug("test text", testArray)
	lg.Debugln("test text", testArray)
	lg.Debugf("some values: %s, %v", "test text", testArray)

	// Trace
	lg.Trace("test text", testArray)
	lg.Traceln("test text", testArray)
	lg.Tracef("some values: %s, %v", "test text", testArray)

	// Warning
	lg.Warn("test text", testArray)
	lg.Warnln("test text", testArray)
	lg.Warnf("some values: %s, %v", "test text", testArray)

	// Error
	lg.Error("test text", testArray)
	lg.Errorln("test text", testArray)
	lg.Errorf("some values: %s, %v", "test text", testArray)

	// Fatal
	lg.Fatal("test text", testArray)
	lg.Fatalln("test text", testArray)
	lg.Fatalf("some values: %s, %v", "test text", testArray)

	// Panic
	lg.Panic("test text", testArray)
	lg.Panicln("test text", testArray)
	lg.Panicf("some values: %s, %v", "test text", testArray)
}
```

### You can turn off the color of logs

For example:  

```go
package main

import "github.com/bearatol/lg"

func main() {
    lg.GetLogger().OffColor()

    lg.Info("Hello, world!")
}

//...
```

or

```go
package main

import "github.com/bearatol/lg"

func init() {
    lg.GetLogger().OffColor()
}

func main() {
    lg.Info("Hello, world!")
}

//...
```

[doc-img]: https://pkg.go.dev/badge/github.com/bearatol/lg
[doc]: https://pkg.go.dev/github.com/bearatol/lg
[ci-img]: https://github.com/bearatol/lg/actions/workflows/lg.yml/badge.svg
[ci]: https://github.com/bearatol/lg/actions/workflows/lg.yml
[report-img]: https://goreportcard.com/badge/github.com/bearatol/lg
[report]: https://goreportcard.com/report/github.com/bearatol/lg
[example-img]: https://user-images.githubusercontent.com/54537638/215755109-8b1b1a3a-4b54-4796-a868-c9b6613f72c9.png
