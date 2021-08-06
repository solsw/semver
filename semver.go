package semver

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	ErrMalformedSemVer = errors.New("malformed semver")
)

// SemVer represents Semantic Versioning Specification (https://semver.org/#semantic-versioning-specification-semver).
type SemVer struct {
	// https://semver.org/#spec-item-8
	Major int
	// https://semver.org/#spec-item-7
	Minor int
	// https://semver.org/#spec-item-6
	Patch int
	// https://semver.org/#spec-item-9
	PreRelease string
	// https://semver.org/#spec-item-10
	Build string
}

// String implements the fmt.Stringer interface.
func (v *SemVer) String() string {
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
func (v *SemVer) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (v *SemVer) UnmarshalText(text []byte) error {
	ss := strings.SplitN(string(text), ".", 3)
	if len(ss) < 3 {
		return ErrMalformedSemVer
	}
	var err error
	if v.Major, err = stringToVersion(ss[0]); err != nil {
		return err
	}
	if v.Minor, err = stringToVersion(ss[1]); err != nil {
		return err
	}
	if err = v.ss2ToPatchPreReleaseBuild(ss[2]); err != nil {
		return err
	}
	return nil
}

func stringToVersion(s string) (int, error) {
	if len(s) == 0 {
		// https://semver.org/#spec-item-2
		return 0, ErrMalformedSemVer
	}
	if s == "0" {
		return 0, nil
	}
	if strings.HasPrefix(s, "0") {
		// https://semver.org/#spec-item-2
		return 0, ErrMalformedSemVer
	}
	ver, err := strconv.Atoi(s)
	if err != nil {
		// https://semver.org/#spec-item-2
		return 0, err
	}
	if ver < 0 {
		// https://semver.org/#spec-item-2
		return 0, ErrMalformedSemVer
	}
	return ver, nil
}

func (v *SemVer) ss2ToPatchPreReleaseBuild(ss2 string) error {
	extIdx := strings.IndexAny(ss2, "-+")
	if extIdx == 0 {
		// no Patch
		return ErrMalformedSemVer
	}
	var err error
	var patch string
	if extIdx == -1 {
		// no extension
		patch = ss2
	} else {
		patch = ss2[:extIdx]
		if err = v.extToPreReleaseBuild(ss2[extIdx+1:], ss2[extIdx] == '+'); err != nil {
			return err
		}
	}
	if v.Patch, err = stringToVersion(patch); err != nil {
		return err
	}
	return nil
}

func (v *SemVer) extToPreReleaseBuild(ext string, noPreRelease bool) error {
	if len(ext) == 0 {
		return ErrMalformedSemVer
	}
	if noPreRelease {
		v.Build = ext
	} else {
		ee := strings.SplitN(ext, "+", 2)
		if len(ee[0]) == 0 {
			return ErrMalformedSemVer
		}
		v.PreRelease = ee[0]
		if len(ee) > 1 {
			if len(ee[1]) == 0 {
				return ErrMalformedSemVer
			}
			v.Build = ee[1]
		}
	}
	if err := validExt(v.PreRelease, true); err != nil {
		return err
	}
	if err := validExt(v.Build, false); err != nil {
		return err
	}
	return nil
}

func validExt(ext string, preRelease bool) error {
	if len(ext) == 0 {
		return nil
	}
	for _, ident := range strings.Split(ext, ".") {
		if err := validIdent(ident, preRelease); err != nil {
			return err
		}
	}
	return nil
}

func validIdent(ident string, preRelease bool) error {
	// https://semver.org/#spec-item-9
	// https://semver.org/#spec-item-10
	if len(ident) == 0 {
		return ErrMalformedSemVer
	}
	if ident == "0" {
		return nil
	}
	for _, r := range ident {
		if !(('0' <= r && r <= '9') || ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z') || r == '-') {
			return ErrMalformedSemVer
		}
	}
	if preRelease {
		if _, err := strconv.Atoi(ident); err == nil {
			// numeric identifier
			if strings.HasPrefix(ident, "0") {
				return ErrMalformedSemVer
			}
		}
	}
	return nil
}

// Valid reports whether 'v' is valid according to https://semver.org/#semantic-versioning-specification-semver.
func (v *SemVer) Valid() bool {
	return v.Major >= 0 && v.Minor >= 0 && v.Patch >= 0 &&
		validExt(v.PreRelease, true) == nil && validExt(v.Build, false) == nil
}

func boolToCompareResult(b bool) int {
	if b {
		return 1
	}
	return -1
}

func comparePreRelease(pr1, pr2 string) int {
	if len(pr1) == 0 && len(pr2) == 0 {
		return 0
	}
	if len(pr1) > 0 && len(pr2) == 0 {
		return -1
	}
	if len(pr1) == 0 && len(pr2) > 0 {
		return 1
	}
	ids1 := strings.Split(pr1, ".")
	ids2 := strings.Split(pr2, ".")
	for i := 0; i < int(math.Min(float64(len(ids1)), float64(len(ids2)))); i++ {
		n1, er1 := strconv.Atoi(ids1[i])
		n2, er2 := strconv.Atoi(ids2[i])
		if er1 == nil && er2 == nil && n1 != n2 {
			return boolToCompareResult(n1 > n2)
		}
		if er1 == nil && er2 != nil {
			return -1
		}
		if er1 != nil && er2 == nil {
			return 1
		}
		if ids1[i] != ids2[i] {
			return boolToCompareResult(ids1[i] > ids2[i])
		}
	}
	if len(ids1) < len(ids2) {
		return -1
	}
	if len(ids1) > len(ids2) {
		return 1
	}
	return 0
}

// Compare compares 'v' with 'other' (see https://semver.org/#spec-item-11).
// Compare returns -1 if 'v' is less than 'other', 0 if 'v' is equal to 'other', 1 if 'v' is greater than 'other'.
func (v *SemVer) Compare(other *SemVer) (int, error) {
	if !(v.Valid() && other.Valid()) {
		return 0, ErrMalformedSemVer
	}
	// https://semver.org/#spec-item-11
	if v.Major != other.Major {
		return boolToCompareResult(v.Major > other.Major), nil
	}
	if v.Minor != other.Minor {
		return boolToCompareResult(v.Minor > other.Minor), nil
	}
	if v.Patch != other.Patch {
		return boolToCompareResult(v.Patch > other.Patch), nil
	}
	return comparePreRelease(v.PreRelease, other.PreRelease), nil
}
