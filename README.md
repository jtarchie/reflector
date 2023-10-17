# reflector

## Description

The `reflector` package provides additional functionality for working with
reflection in Go. It includes types and functions that make it easier to access
fields and tags in structs using reflection.

## Features

- The `Type` type provides a wrapper around `reflect.Type` and includes methods
  for accessing fields.
- The `StructField` type wraps around `reflect.StructField` and provides a way
  to access field tags.
- The `FieldByName` function retrieves a field by its name from a struct type.
- The `GetTag` function returns the value of a specific tag for a field.

## Usage

To use the `reflector` package, you can import it into your Go code:

```bash
go get github.com/jtarchie/reflector
```

### Example

```go
package main

import (
	"fmt"
	"github.com/example/reflector"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	t := reflector.TypeOf(User{})
	if field, ok := t.FieldByName("Name"); ok {
		fmt.Println("Field name:", field.Name)
		fmt.Println("Field tag:", field.GetTag("json"))
	}
}
```

This example demonstrates how to use the `reflector` package to access a field's
name and tag in a struct. The `TypeOf` function is used to get the reflection
type of the `User` struct, and then the `FieldByName` function retrieves the
field with the given name. The `GetTag` function is used to access the value of
the `json` tag for the field.

## Benchmarking

The `benchmark_test.go` file contains benchmark tests for the `reflector`
package. These tests compare the performance of different methods for accessing
a field's type and JSON tag using the reflect and reflector packages.

To run the benchmark tests, use the following command:

```
go test -bench=.
```

There are four benchmark functions provided:

- `BenchmarkControl` measures the performance of accessing a field's type using
  the `reflect` package.
- `BenchmarkCacheField` measures the performance of accessing a field's type
  using the `reflector` package with field caching.
- `BenchmarkCacheTypeTag` measures the performance of accessing a field's type
  and tag using the `reflector` package with type and tag caching.
- `BenchmarkReflector` measures the performance of accessing a field's type and
  tag using the `reflector` package without caching.

```
BenchmarkControl-8              17058214                59.08 ns/op
BenchmarkCacheField-8           19952377                59.16 ns/op
BenchmarkCacheTypeTag-8         46129893                24.18 ns/op
BenchmarkReflector-8            42407883                27.77 ns/op
```
