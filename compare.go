package semver

import (
	"strings"
)

// Compare [compares] 'sv1' with 'sv2'.
// Compare returns -1 if 'sv1' is less than 'sv2', 0 if 'sv1' is equal to 'sv2', 1 if 'sv1' is more than 'sv2'.
//
// [compares]: https://semver.org/#spec-item-11
func Compare(sv1, sv2 SemVer) (int, error) {
	if err := Valid(sv1); err != nil {
		return 0, err
	}
	if err := Valid(sv2); err != nil {
		return 0, err
	}
	// https://semver.org/#spec-item-11
	if sv1.Major != sv2.Major {
		return boolToCompareResult(sv1.Major > sv2.Major), nil
	}
	if sv1.Minor != sv2.Minor {
		return boolToCompareResult(sv1.Minor > sv2.Minor), nil
	}
	if sv1.Patch != sv2.Patch {
		return boolToCompareResult(sv1.Patch > sv2.Patch), nil
	}
	return comparePreRelease(sv1.PreRelease, sv2.PreRelease), nil
}

// Less reports whether 'sv1' is [less] than 'sv2'.
// If 'sv1' or/and 'sv2' is/are invalid, Less panics.
//
// [less]: https://semver.org/#spec-item-11
func Less(sv1, sv2 SemVer) bool {
	r, err := Compare(sv1, sv2)
	if err != nil {
		panic(err)
	}
	return r < 0
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
	shorter := len(ids1)
	if len(ids2) < shorter {
		shorter = len(ids2)
	}
	for i := 0; i < shorter; i++ {
		id1, id2 := ids1[i], ids2[i]
		num1, num2 := isNumeric(id1), isNumeric(id2)
		switch {
		case num1 && num2:
			// https://semver.org/#spec-item-11 (4.1): compare numerically.
			// Identifiers are validated to have no leading zeros, so the
			// longer digit string is the larger number; equal lengths
			// compare lexically. This avoids integer overflow.
			if len(id1) != len(id2) {
				return boolToCompareResult(len(id1) > len(id2))
			}
			if id1 != id2 {
				return boolToCompareResult(id1 > id2)
			}
		case num1 && !num2:
			// https://semver.org/#spec-item-11 (4.3)
			return -1
		case !num1 && num2:
			return 1
		default:
			// https://semver.org/#spec-item-11 (4.2): both alphanumeric.
			if id1 != id2 {
				return boolToCompareResult(id1 > id2)
			}
		}
	}
	// https://semver.org/#spec-item-11 (4.4)
	if len(ids1) < len(ids2) {
		return -1
	}
	if len(ids1) > len(ids2) {
		return 1
	}
	return 0
}
