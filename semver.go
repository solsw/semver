package semver

import (
	"errors"
	"fmt"
)

var (
	ErrMalformedSemVer = errors.New("malformed semver")
	ErrNoSemVer        = errors.New("no semver")
)

// SemVer represents Semantic Versioning Specification
// (https://semver.org/#semantic-versioning-specification-semver).
// SemVer's zero value is "0.0.0" version.
type SemVer struct {
	// major version (https://semver.org/#spec-item-8)
	Major int
	// minor version (https://semver.org/#spec-item-7)
	Minor int
	// patch version (https://semver.org/#spec-item-6)
	Patch int
	// pre-release version (https://semver.org/#spec-item-9)
	PreRelease string
	// build metadata (https://semver.org/#spec-item-10)
	Build string
}

// String implements the fmt.Stringer interface.
func (v SemVer) String() string {
	s := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if len(v.PreRelease) > 0 {
		s += "-" + v.PreRelease
	}
	if len(v.Build) > 0 {
		s += "+" + v.Build
	}
	return s
}

// MarshalText implements the encoding.TextMarshaler interface.
func (v SemVer) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (v *SemVer) UnmarshalText(text []byte) error {
	sv, err := Parse(string(text))
	if err != nil {
		return err
	}
	*v = sv
	return nil
}

// IsValid reports whether 'v' is valid (https://semver.org/#semantic-versioning-specification-semver).
func (v SemVer) IsValid() bool {
	return Valid(v)
}

// CompareTo compares 'v' with 'other' (https://semver.org/#spec-item-11).
// CompareTo returns -1 if 'v' is less than 'other', 0 if 'v' is equal to 'other', 1 if 'v' is greater than 'other'.
func (v SemVer) CompareTo(other SemVer) (int, error) {
	return Compare(v, other)
}
