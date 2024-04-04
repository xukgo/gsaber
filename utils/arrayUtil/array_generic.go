package arrayUtil

// IndexOfArray 返回切片中第一次出现目标值的索引位置。
// 如果找不到，则返回-1。
// 这是一个泛型函数，支持任何可比较的类型。
func IndexOfArray[T comparable](s []T, target T) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}
	return -1
}
func SplitIntoChunks[T any](s []T, chunkSize int) [][]T {
	var chunks [][]T
	for chunkSize < len(s) {
		s, chunks = s[chunkSize:], append(chunks, s[0:chunkSize:chunkSize])
	}
	chunks = append(chunks, s)
	return chunks
}

func SpreadNilToCountList[T any](items []T, count int) []T {
	if len(items) == 0 {
		return make([]T, 0)
	}

	minCount := len(items)
	if len(items) > count {
		minCount = count
	}

	res := make([]T, count)
	for i := 0; i < minCount; i++ {
		res[i] = items[i]
	}
	return res
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
