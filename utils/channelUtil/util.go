package channelUtil

func TrySend[T any](ch chan T, value T) bool {
	select {
	case ch <- value:
		return true
	default:
		return false
	}
}

func TryRecv[T any](ch chan T, value T) (v T, br bool) {
	select {
	case v = <-ch:
		return v, true
	default:
		return v, false
	}
}
