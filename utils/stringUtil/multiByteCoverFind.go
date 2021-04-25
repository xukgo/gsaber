package stringUtil

type MultiByteCover struct {
	Same  bool
	Left  []byte
	Right []byte
}

func checkBytesEqual(left, right []byte) bool {
	if len(left) != len(right) {
		return false
	}
	for idx := range left {
		if left[idx] != right[idx] {
			return false
		}
	}
	return true
}

//target为比较的标准
func checkLenBytesEqual(target, buff []byte) bool {
	if len(buff) < len(target) {
		return false
	}
	for idx := range target {
		if target[idx] != buff[idx] {
			return false
		}
	}
	return true
}

func InitMultiByteCover(left, right []byte) MultiByteCover {
	return MultiByteCover{
		Left:  left,
		Right: right,
		Same:  checkBytesEqual(left, right),
	}
}

type MultiByteCoverFinder struct {
	Covers []MultiByteCover
}

func (this *MultiByteCoverFinder) AddCover(cover MultiByteCover) {
	if len(this.Covers) == 0 {
		this.Covers = append(this.Covers, cover)
		return
	}
	index := -1
	for idx := range this.Covers {
		if len(this.Covers[idx].Left) < len(cover.Left) {
			index = idx
			break
		}
	}
	this.Covers = append(this.Covers, cover)
	if index < 0 {
		return
	}
	for i := len(this.Covers) - 2; i >= index; i-- {
		this.Covers[i+1] = this.Covers[i]
	}
	this.Covers[index] = cover
}

func (this *MultiByteCoverFinder) GetSegments(buffer []byte, buffLen int, left []byte, right []byte) [][]byte {
	resp := make([][]byte, 0, 1)
	for {
		lindex, rindex := this.getSegmentIndex(buffer, buffLen, left, right)
		if lindex >= 0 && rindex > 0 {
			resp = append(resp, buffer[lindex:rindex+len(right)])
		}
		if rindex+len(right) >= buffLen {
			break
		}
		buffer = buffer[rindex+len(right):]
		buffLen = buffLen - rindex - len(right)
	}
	return resp
}
func (this *MultiByteCoverFinder) GetFirstSegment(buffer []byte, buffLen int, left []byte, right []byte) []byte {
	lindex, rindex := this.getSegmentIndex(buffer, buffLen, left, right)
	if lindex >= 0 && rindex > 0 {
		return buffer[lindex : rindex+len(right)]
	}
	return nil
}
func (this *MultiByteCoverFinder) getSegmentIndex(buffer []byte, buffLen int, left []byte, right []byte) (int, int) {
	leftIndex := this.Index(buffer, buffLen, left)
	if leftIndex < 0 {
		return -1, -1
	}
	if leftIndex+len(left) >= buffLen {
		return leftIndex, -1
	}
	buffer2 := buffer[leftIndex+len(left):]
	RightIndex := this.Index(buffer2, buffLen-leftIndex-len(left), right)
	if RightIndex < 0 {
		return leftIndex, -1
	}
	return leftIndex, RightIndex + leftIndex + len(left)
}
func (this *MultiByteCoverFinder) Index(buffer []byte, buffLen int, target []byte) int {
	coverPairs := make([]intPair, 0, len(this.Covers))
	for idx := 0; idx < buffLen; {
		v := buffer[idx:]
		if len(coverPairs) == 0 && checkLenBytesEqual(target, v) {
			return idx
		}

		nowIndex, nowDirect := this.compareByteCovers(v)
		if nowIndex < 0 {
			idx++
			continue
		}
		var nowLen = 0
		if nowDirect == 1 {
			nowLen = len(this.Covers[nowIndex].Left)
		} else {
			nowLen = len(this.Covers[nowIndex].Right)
		}

		if len(coverPairs) == 0 {
			if this.Covers[nowIndex].Same {
				coverPairs = append(coverPairs, intPair{nowIndex, 1})
				idx += nowLen
				continue
			}
			if nowDirect == 2 {
				idx++
				continue
			}
			coverPairs = append(coverPairs, intPair{nowIndex, 1})
			idx += nowLen
			continue
		}

		currentCoverIndex := coverPairs[len(coverPairs)-1].Index
		//不匹配 看是否往后追加cover
		if currentCoverIndex != nowIndex {
			if nowDirect == 1 {
				coverPairs = append(coverPairs, intPair{nowIndex, 1})
				idx += nowLen
				continue
			}
			idx++
			continue
		} else {
			nowCover := this.Covers[nowIndex]
			if nowCover.Same {
				coverPairs = deleteTail(coverPairs)
				idx += nowLen
				continue
			}
			if nowDirect == 2 {
				coverPairs = deleteTail(coverPairs)
				idx += nowLen
				continue
			}

			coverPairs = append(coverPairs, intPair{nowIndex, 1})
			idx += nowLen
			continue
		}
	}
	return -1
}

//compareByteCovers 0不匹配 1匹配左 2匹配右
func (this MultiByteCoverFinder) compareByteCovers(v []byte) (int, int) {
	for idx := range this.Covers {
		if checkLenBytesEqual(this.Covers[idx].Left, v) {
			return idx, 1
		}
		if checkLenBytesEqual(this.Covers[idx].Right, v) {
			return idx, 2
		}
		continue
	}
	return -1, 0
}
