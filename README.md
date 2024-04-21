[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/pythonik/golang-commons)](https://goreportcard.com/report/github.com/pythonik/golang-commons)
[![codecov](https://codecov.io/gh/pythonik/golang-commons/graph/badge.svg?token=YSMFAR20IF)](https://codecov.io/gh/pythonik/golang-commons)
[![Unit test](https://github.com/pythonik/golang-commons/actions/workflows/ci.yml/badge.svg)](https://github.com/pythonik/golang-commons/actions/workflows/ci.yml)
# golang-commons 
This project is a collection of common utilities and libraries for Golang inspired by Google Guava and Apache Commons Lang. It is a work in progress and will be updated as new utilities are added.

## Installation

To install this library, use the following `go get` command:

```bash
go get -u github.com/pythonik/golang-commons
```

## Usage

### Importing the package accordingly

the example below uses `collections`:

```go
import (
    "github.com/pythonik/golang-commons/collections"
)
```

### Creating a New BiMap

You can create a new instance of a `BiMap` using the `NewBiMap` function. Here's an example for a map with `string` keys and `int` values:

```go
bimap := collections.NewBiMap[string, int]()
```

### Adding Entries to BiMap

To add entries to the `BiMap`, use the `Put` method:

```go
bimap.Put("apple", 1)
bimap.Put("banana", 2)
```

### Retrieving Values by Key

To retrieve a value based on a key, use the `GetByKey` method:

```go
value, ok := bimap.GetByKey("apple")
if ok {
    fmt.Println("The value for 'apple' is:", value)
}
```

### Retrieving Keys by Value

To retrieve a key based on a value, use the `GetByValue` method:

```go
key, ok := bimap.GetByValue(2)
if ok {
    fmt.Println("The key for value '2' is:", key)
}
```