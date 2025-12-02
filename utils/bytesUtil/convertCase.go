package bytesUtil

func ToUpperAsciiSelf(s []byte) {
	var sub byte = 'a' - 'A'
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'a' <= c && c <= 'z' {
			c -= sub
		}
		s[i] = c
	}
}

func ToLowerAsciiSelf(s []byte) {
	var sub byte = 'a' - 'A'
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c += sub
		}
		s[i] = c
	}
}
