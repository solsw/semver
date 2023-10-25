package semver

import (
	"fmt"
)

// SemVer represents [Semantic Versioning Specification].
// SemVer's [zero value] is "0.0.0" version.
//
// [Semantic Versioning Specification]: https://semver.org/#semantic-versioning-specification-semver
// [zero value]: https://go.dev/ref/spec#The_zero_value
type SemVer struct {
	// Major contains SemVer's [major version].
	//
	// [major version]: https://semver.org/#spec-item-8
	Major int
	// Minor contains SemVer's [minor version].
	//
	// [minor version]: https://semver.org/#spec-item-7
	Minor int
	// Patch contains SemVer's [patch version].
	//
	// [patch version]: https://semver.org/#spec-item-6
	Patch int
	// PreRelease contains SemVer's [pre-release version].
	//
	// [pre-release version]: https://semver.org/#spec-item-9
	PreRelease string
	// Build contains SemVer's [build metadata].
	//
	// [build metadata]: https://semver.org/#spec-item-10
	Build string
}

// String implements the [fmt.Stringer] interface.
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

// MarshalText implements the [encoding.TextMarshaler] interface.
func (v SemVer) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (v *SemVer) UnmarshalText(text []byte) error {
	sv, err := Parse(string(text))
	if err != nil {
		return err
	}
	*v = sv
	return nil
}

// IsValid reports whether 'v' is [valid].
//
// [valid]: https://semver.org/#semantic-versioning-specification-semver
func (v SemVer) IsValid() bool {
	return Valid(v) == nil
}

// CompareTo [compares] 'v' with 'other'.
// CompareTo returns -1 if 'v' is less than 'other', 0 if 'v' is equal to 'other', 1 if 'v' is more than 'other'.
//
// [compares]: https://semver.org/#spec-item-11
func (v SemVer) CompareTo(other SemVer) (int, error) {
	return Compare(v, other)
}

// LessThan reports whether 'v' is [less] than 'other'.
//
// [less]: https://semver.org/#spec-item-11
func (v SemVer) LessThan(other SemVer) (bool, error) {
	r, err := Compare(v, other)
	return r < 0, err
}

// EqualTo reports whether 'v' is [equal] to 'other'.
//
// [equal]: https://semver.org/#spec-item-11
func (v SemVer) EqualTo(other SemVer) (bool, error) {
	r, err := Compare(v, other)
	return r == 0, err
}

// MoreThan reports whether 'v' is [more] than 'other'.
//
// [more]: https://semver.org/#spec-item-11
func (v SemVer) MoreThan(other SemVer) (bool, error) {
	r, err := Compare(v, other)
	return r > 0, err
}
