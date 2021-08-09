package semver

import (
	"strconv"
	"strings"
)

// Valid reports whether 'sv' is valid according to https://semver.org/#semantic-versioning-specification-semver.
func Valid(sv *SemVer) bool {
	return sv.Major >= 0 && sv.Minor >= 0 && sv.Patch >= 0 &&
		validExt(sv.PreRelease, true) == nil && validExt(sv.Build, false) == nil
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