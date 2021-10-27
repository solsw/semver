package semver

import (
	"strconv"
	"strings"
)

// Parse parses version string.
func Parse(s string) (*SemVer, error) {
	ss := strings.SplitN(s, ".", 3)
	if len(ss) < 3 {
		return nil, ErrMalformedSemVer
	}
	var sv SemVer
	var err error
	if sv.Major, err = strToVersionNumber(ss[0]); err != nil {
		return nil, err
	}
	if sv.Minor, err = strToVersionNumber(ss[1]); err != nil {
		return nil, err
	}
	if err = ss2ToPatchPreReleaseBuild(&sv, ss[2]); err != nil {
		return nil, err
	}
	return &sv, nil
}

// ParseMust parses version string.
// ParseMust panics in case of parsing error.
func ParseMust(s string) *SemVer {
	sv, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return sv
}

func strToVersionNumber(s string) (int, error) {
	if len(s) == 0 {
		return 0, ErrMalformedSemVer
	}
	// https://semver.org/#spec-item-2
	if s == "0" {
		return 0, nil
	}
	if strings.HasPrefix(s, "0") {
		return 0, ErrMalformedSemVer
	}
	ver, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	if ver < 0 {
		return 0, ErrMalformedSemVer
	}
	return ver, nil
}

func ss2ToPatchPreReleaseBuild(sv *SemVer, ss2 string) error {
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
		if err = extToPreReleaseBuild(sv, ss2[extIdx+1:], ss2[extIdx] == '+'); err != nil {
			return err
		}
	}
	if sv.Patch, err = strToVersionNumber(patch); err != nil {
		return err
	}
	return nil
}

func extToPreReleaseBuild(sv *SemVer, ext string, noPreRelease bool) error {
	if len(ext) == 0 {
		return ErrMalformedSemVer
	}
	if noPreRelease {
		sv.Build = ext
	} else {
		ee := strings.SplitN(ext, "+", 2)
		if len(ee[0]) == 0 {
			return ErrMalformedSemVer
		}
		sv.PreRelease = ee[0]
		if len(ee) > 1 {
			if len(ee[1]) == 0 {
				return ErrMalformedSemVer
			}
			sv.Build = ee[1]
		}
	}
	if err := validExt(sv.PreRelease, true); err != nil {
		return err
	}
	if err := validExt(sv.Build, false); err != nil {
		return err
	}
	return nil
}
