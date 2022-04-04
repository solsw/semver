package semver

import (
	"errors"
	"strconv"
	"strings"
)

// Valid checks 'sv' for validity (https://semver.org/#semantic-versioning-specification-semver).
// If 'sv' is not valid corresponding error is reurned.
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
	if preRelease {
		if _, err := strconv.Atoi(ident); err == nil {
			// numeric identifier
			if strings.HasPrefix(ident, "0") {
				return errors.New("malformed semver")
			}
		}
	}
	return nil
}
