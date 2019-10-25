package fileWatcher

type Observer interface {
	UpdateFromContent(content []byte)
}
