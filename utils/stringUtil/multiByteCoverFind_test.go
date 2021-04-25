package stringUtil

import (
	"testing"
)

func Test_multiByteAddCover(t *testing.T) {
	finder := new(MultiByteCoverFinder)
	finder.AddCover(InitMultiByteCover([]byte{'(', '('}, []byte{')', ')'}))
	finder.AddCover(InitMultiByteCover([]byte{'(', '(', '('}, []byte{')', ')', ')'}))
	finder.AddCover(InitMultiByteCover([]byte{'('}, []byte{')'}))
	if len(finder.Covers) != 3 {
		t.FailNow()
	}
	if len(finder.Covers[0].Left) != 3 {
		t.FailNow()
	}
	if len(finder.Covers[1].Left) != 2 {
		t.FailNow()
	}
	if len(finder.Covers[2].Left) != 1 {
		t.FailNow()
	}
}
func Test_multiByteCoverFind1(t *testing.T) {
	str := "`rate` decimal((11,4)) DEFAULT NULL,,"
	finder := new(MultiByteCoverFinder)
	finder.AddCover(InitMultiByteCover([]byte{'(', '('}, []byte{')', ')'}))
	finder.AddCover(InitMultiByteCover([]byte{'`'}, []byte{'`'}))
	buff := []byte(str)
	index := finder.Index(buff, len(buff), []byte{',', ','})
	if index != len(buff)-2 {
		t.FailNow()
	}
	index = finder.Index(buff, len(buff), []byte{')'})
	if index != -1 {
		t.FailNow()
	}
}
func Test_multiByteCoverFind2(t *testing.T) {
	str := "1+(d+'var1*(b*c)')*(a+(b*c/(d+f)))@"
	finder := new(MultiByteCoverFinder)
	finder.AddCover(InitMultiByteCover([]byte{'('}, []byte{')'}))
	finder.AddCover(InitMultiByteCover([]byte{0x27}, []byte{0x27}))
	buff := []byte(str)
	index := finder.Index(buff, len(buff), []byte{'@'})
	if index != len(buff)-1 {
		t.FailNow()
	}
	index = finder.Index(buff, len(buff), []byte{')'})
	if index != -1 {
		t.FailNow()
	}
}

func Test_multiByteGetFirstCompareSegment(t *testing.T) {
	str := "1+(d+'var1*(b*c)')*(a+(b*c/(d+f)))"
	finder := new(MultiByteCoverFinder)
	finder.AddCover(InitMultiByteCover([]byte{'('}, []byte{')'}))
	finder.AddCover(InitMultiByteCover([]byte{0x27}, []byte{0x27}))
	buff := []byte(str)
	seg := finder.GetFirstSegment(buff, len(buff), []byte{'('}, []byte{')'})
	if len(seg) != 16 {
		t.FailNow()
	}
	segs := finder.GetSegments(buff, len(buff), []byte{'('}, []byte{')'})
	if len(segs) != 2 {
		t.FailNow()
	}
}
