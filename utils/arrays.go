package utils

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CompareStringArrays(a, b []string) (aOnly []string, bOnly []string, both []string, equal bool) {
	for _, v := range a {
		if !ContainsString(b, v) {
			aOnly = append(aOnly, v)
		} else {
			both = append(both, v)
		}
	}
	for _, v := range b {
		if !ContainsString(a, v) {
			bOnly = append(bOnly, v)
		}
	}
	return aOnly, bOnly, both, len(aOnly) == 0 && len(bOnly) == 0
}
