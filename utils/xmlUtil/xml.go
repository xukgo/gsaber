package xmlUtil

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"regexp"
	"unicode/utf8"
)

func ReadCleanXml(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	reg := regexp.MustCompile(`<!--.*?-->`)
	//reg := regexp.MustCompile(`@@(.*?)@@`)
	subs := reg.FindAll(content, -1)

	for idx := range subs {
		content = bytes.ReplaceAll(content, subs[idx], []byte(" "))
	}
	return content, nil
}

//func hanlderLineReadFile(fileName string, handler func(string)) error {
//	f, err := os.Open(fileName)
//	if err != nil {
//		return err
//	}
//	buf := bufio.NewReader(f)
//	for {
//		line, err := buf.ReadString('\n')
//		line = strings.TrimSpace(line)
//		handler(line)
//		if err != nil {
//			if err == io.EOF {
//				return nil
//			}
//			return err
//		}
//	}
//}

//e.Indent("", "\t")
func Pretty(src []byte, prefix, indent string) []byte {
	var b bytes.Buffer
	br := bytes.NewReader(src)
	bw := bufio.NewWriter(&b)
	d := xml.NewDecoder(br)
	e := xml.NewEncoder(bw)
	e.Indent(prefix, indent)

	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}

		if tok, ok := t.(xml.CharData); ok {
			r, _ := utf8.DecodeRune(tok)
			if whitespace(r) || endOfLine(r) {
				continue
			}
		}

		e.EncodeToken(t)
	}
	e.Flush()
	return b.Bytes()
}

func endOfLine(ch rune) bool {
	return ch == '\n' || ch == '\r'
}

func whitespace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}
