package slicer

type Comparable interface {
	IsEqual(obj interface{}) bool
}
