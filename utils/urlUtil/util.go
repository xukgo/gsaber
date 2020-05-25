package urlUtil

import (
	"bytes"
)

func Combind(url, tail string) string {
	for {
		if url[len(url)-1] == '/' {
			url = url[:len(url)-1]
			continue
		}
		break
	}
	for {
		if tail[0] == '/' {
			tail = tail[1:]
			continue
		}
		break
	}

	bf := new(bytes.Buffer)
	bf.WriteString(url)
	bf.WriteString("/")
	bf.WriteString(tail)
	return bf.String()
}
