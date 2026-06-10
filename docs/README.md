# semver
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/semver.svg)](https://pkg.go.dev/github.com/solsw/semver)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/semver)

Package **semver** contains [Semantic Versioning 2.0.0](https://semver.org/) support for Go.

## Installation

```sh
go get github.com/solsw/semver
```

## SemVer

A version is represented by the `SemVer` struct:

```go
type SemVer struct {
	Major      int64  // https://semver.org/#spec-item-8
	Minor      int64  // https://semver.org/#spec-item-7
	Patch      int64  // https://semver.org/#spec-item-6
	PreRelease string // https://semver.org/#spec-item-9
	Build      string // https://semver.org/#spec-item-10
}
```

The zero value of `SemVer` is the `"0.0.0"` version.

`SemVer` implements [`fmt.Stringer`](https://pkg.go.dev/fmt#Stringer),
[`encoding.TextMarshaler`](https://pkg.go.dev/encoding#TextMarshaler) and
[`encoding.TextUnmarshaler`](https://pkg.go.dev/encoding#TextUnmarshaler),
so it can be used directly with `fmt` and with JSON, XML, etc.

> `String` and `MarshalText` do not validate the receiver. An invalid `SemVer`
> (e.g. with negative fields) may produce an invalid version string. Use
> `IsValid` to check it beforehand.

## Usage

```go
package main

import (
	"fmt"

	"github.com/solsw/semver"
)

func main() {
	v, err := semver.Parse("1.2.3-rc.1+build.7")
	if err != nil {
		panic(err)
	}
	fmt.Println(v.Major, v.Minor, v.Patch) // 1 2 3
	fmt.Println(v.PreRelease)              // rc.1
	fmt.Println(v.Build)                   // build.7
	fmt.Println(v)                         // 1.2.3-rc.1+build.7

	less, err := v.LessThan(semver.SemVer{Major: 1, Minor: 2, Patch: 3})
	if err != nil {
		panic(err)
	}
	fmt.Println(less) // true (a pre-release is lower than the release)
}
```

## API

### Functions

- `Parse(s string) (SemVer, error)` — convert a version string to a `SemVer`.
- `Valid(sv SemVer) error` — report whether `sv` is valid (returns the corresponding error otherwise).
- `Compare(sv1, sv2 SemVer) (int, error)` — compare two versions; returns `-1`, `0` or `1`.
- `Less(sv1, sv2 SemVer) bool` — report whether `sv1` is less than `sv2` (panics on an invalid version).

### Methods

- `(v SemVer) String() string`
- `(v SemVer) MarshalText() ([]byte, error)`
- `(v *SemVer) UnmarshalText(text []byte) error`
- `(v SemVer) IsValid() bool`
- `(v SemVer) CompareTo(other SemVer) (int, error)`
- `(v SemVer) LessThan(other SemVer) (bool, error)`
- `(v SemVer) EqualTo(other SemVer) (bool, error)`
- `(v SemVer) MoreThan(other SemVer) (bool, error)`
