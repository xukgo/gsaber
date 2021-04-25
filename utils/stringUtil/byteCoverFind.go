package stringUtil

// ByteCover
//英文单引号 ' 0x27
//英文双引号 " 0x21
//重音符 ` 0x60
//\n 0x0A 换行(LF) ，将当前位置移到下一行开头
//\r 0x0d 回车(CR) ，将当前位置移到本行开头
//\t 0x09 水平制表(HT) （跳到下一个TAB位置）

type ByteCover struct {
	Same  bool
	Left  byte
	Right byte
}

func InitByteCover(left, right byte) ByteCover {
	return ByteCover{
		Left:  left,
		Right: right,
		Same:  left == right,
	}
}

type intPair struct {
	Index     int //coverIndex
	Direction int //left or right
}

type ByteCoverFinder struct {
	Covers []ByteCover
}

func (this *ByteCoverFinder) AddCover(cover ByteCover) {
	this.Covers = append(this.Covers, cover)
}

func (this *ByteCoverFinder) GetSegments(buffer []byte, buffLen int, left byte, right byte) [][]byte {
	resp := make([][]byte, 0, 1)
	for {
		lindex, rindex := this.getSegmentIndex(buffer, buffLen, left, right)
		if lindex >= 0 && rindex > 0 {
			resp = append(resp, buffer[lindex:rindex+1])
		}
		if rindex >= buffLen-1 {
			break
		}
		buffer = buffer[rindex+1:]
		buffLen = buffLen - rindex - 1
	}
	return resp
}
func (this *ByteCoverFinder) GetFirstSegment(buffer []byte, buffLen int, left byte, right byte) []byte {
	lindex, rindex := this.getSegmentIndex(buffer, buffLen, left, right)
	if lindex >= 0 && rindex > 0 {
		return buffer[lindex : rindex+1]
	}
	return nil
}
func (this *ByteCoverFinder) getSegmentIndex(buffer []byte, buffLen int, left byte, right byte) (int, int) {
	leftIndex := this.Index(buffer, buffLen, left)
	if leftIndex < 0 {
		return -1, -1
	}
	if leftIndex >= buffLen-1 {
		return leftIndex, -1
	}
	buffer = buffer[leftIndex:]
	RightIndex := this.Index(buffer[1:], buffLen-1, right)
	if RightIndex < 0 {
		return leftIndex, -1
	}
	RightIndex = RightIndex + 1
	return leftIndex, RightIndex + leftIndex
}

func (this *ByteCoverFinder) Index(buffer []byte, buffLen int, target byte) int {
	coverPairs := make([]intPair, 0, len(this.Covers))
	for idx := 0; idx < buffLen; idx++ {
		v := buffer[idx]
		if len(coverPairs) == 0 && v == target {
			return idx
		}

		nowIndex, nowDirect := this.compareByteCovers(v)
		if nowIndex < 0 {
			continue
		}

		if len(coverPairs) == 0 {
			if this.Covers[nowIndex].Same {
				coverPairs = append(coverPairs, intPair{nowIndex, 1})
				continue
			}
			if nowDirect == 2 {
				continue
			}
			coverPairs = append(coverPairs, intPair{nowIndex, 1})
			continue
		}

		currentCoverIndex := coverPairs[len(coverPairs)-1].Index
		//不匹配 看是否往后追加cover
		if currentCoverIndex != nowIndex {
			if nowDirect == 1 {
				coverPairs = append(coverPairs, intPair{nowIndex, 1})
			}
			continue
		} else {
			nowCover := this.Covers[nowIndex]
			if nowCover.Same {
				coverPairs = deleteTail(coverPairs)
				continue
			}
			if nowDirect == 2 {
				coverPairs = deleteTail(coverPairs)
				continue
			}

			coverPairs = append(coverPairs, intPair{nowIndex, 1})
			continue
		}
	}
	return -1
}

func deleteTail(pairs []intPair) []intPair {
	return pairs[:len(pairs)-1]
}

//compareByteCovers 0不匹配 1匹配左 2匹配右
func (this ByteCoverFinder) compareByteCovers(v byte) (int, int) {
	for idx := range this.Covers {
		if this.Covers[idx].Left == v {
			return idx, 1
		}
		if this.Covers[idx].Right == v {
			return idx, 2
		}
		continue
	}
	return -1, 0
}
