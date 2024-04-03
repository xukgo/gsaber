package arrayUtil

func SplitIntoChunks[T any](s []T, chunkSize int) [][]T {
	var chunks [][]T
	for chunkSize < len(s) {
		s, chunks = s[chunkSize:], append(chunks, s[0:chunkSize:chunkSize])
	}
	chunks = append(chunks, s)
	return chunks
}

func SpreadLoopToCountList[T any](items []T, count int) []T {
	if len(items) == 0 {
		return make([]T, 0)
	}
	res := make([]T, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, GetModIndexItem(items, i))
	}
	return res
}

func GetModIndexItem[T any](items []T, index int) T {
	return items[index%(len(items))]
}
