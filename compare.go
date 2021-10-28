package semver

import (
	"strconv"
	"strings"
)

// Compare compares 'sv1' with 'sv2' (https://semver.org/#spec-item-11).
// Compare returns -1 if 'sv1' is less than 'sv2', 0 if 'sv1' is equal to 'sv2', 1 if 'sv1' is greater than 'sv2'.
func Compare(sv1, sv2 SemVer) (int, error) {
	if !(Valid(sv1) && Valid(sv2)) {
		return 0, ErrMalformedSemVer
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

func boolToCompareResult(b bool) int {
	if b {
		return 1
	}
	return -1
}

func minInt(i, j int) int {
	if i < j {
		return i
	}
	return j
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
	for i := 0; i < minInt(len(ids1), len(ids2)); i++ {
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

// Less reports whether 'sv1' is less than 'sv2' (https://semver.org/#spec-item-11).
// If 'sv1' or/and 'sv2' is/are invalid, Less panics.
// Less is intended for use with "sort" package.
func Less(sv1, sv2 SemVer) bool {
	i, err := Compare(sv1, sv2)
	if err != nil {
		panic(err)
	}
	return i < 0
}
