package stringUtil

var aAGap uint8 = 'a' - 'A'

func CompareIgnoreCase(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}

		if a[i] >= 'a' && a[i] <= 'z' {
			if a[i]-aAGap == b[i] {
				continue
			}
		}

		if a[i] >= 'A' && a[i] <= 'Z' {
			if a[i]+aAGap == b[i] {
				continue
			}
		}

		return false
	}

	return true
}
