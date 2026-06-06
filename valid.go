package semver

import (
	"errors"
	"strings"
)

// isNumeric reports whether s is a non-empty string of ASCII digits only.
// Such strings are "numeric identifiers" per https://semver.org/#spec-item-11.
func isNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// Valid checks 'sv' for [validity]. If 'sv' is not valid corresponding error is returned.
//
// [validity]: https://semver.org/#semantic-versioning-specification-semver
func Valid(sv SemVer) error {
	// https://semver.org/#spec-item-2
	if sv.Major < 0 || sv.Minor < 0 || sv.Patch < 0 {
		return errors.New("malformed semver")
	}
	// https://semver.org/#spec-item-9
	if err := validExt(sv.PreRelease, true); err != nil {
		return err
	}
	// https://semver.org/#spec-item-10
	if err := validExt(sv.Build, false); err != nil {
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
		return errors.New("malformed semver")
	}
	if ident == "0" {
		return nil
	}
	for _, r := range ident {
		if !(('0' <= r && r <= '9') || ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z') || r == '-') {
			return errors.New("malformed semver")
		}
	}
	if preRelease && isNumeric(ident) {
		// numeric identifier must not have leading zeros
		if strings.HasPrefix(ident, "0") {
			return errors.New("malformed semver")
		}
	}
	return nil
}
