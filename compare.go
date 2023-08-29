package semver

import (
	"strconv"
	"strings"
)

// Compare [compares] 'sv1' with 'sv2'.
//
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
//
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
	min := len(ids1)
	if len(ids2) < min {
		min = len(ids2)
	}
	for i := 0; i < min; i++ {
		n1, er1 := strconv.Atoi(ids1[i])
		n2, er2 := strconv.Atoi(ids2[i])
		// https://semver.org/#spec-item-11 (4.1)
		if er1 == nil && er2 == nil && n1 != n2 {
			return boolToCompareResult(n1 > n2)
		}
		// https://semver.org/#spec-item-11 (4.3)
		if er1 == nil && er2 != nil {
			return -1
		}
		if er1 != nil && er2 == nil {
			return 1
		}
		// https://semver.org/#spec-item-11 (4.2)
		if ids1[i] != ids2[i] {
			return boolToCompareResult(ids1[i] > ids2[i])
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
