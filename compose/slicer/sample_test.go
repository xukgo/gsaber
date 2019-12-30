package slicer

import (
	"fmt"
	"testing"
)

func TestSlicer(t *testing.T) {
	intSliceExec()
	fmt.Println("")
	stringSliceExec()
	fmt.Println("")
	structSliceExec()
}

func structSliceExec() {
	fmt.Println("struct slice start")
	xiaoMing := Student{"1001", "xiao ming"}
	xiaoLei := Student{"1002", "xiao lei"}
	xiaoFang := Student{"1003", "xiao fang"}
	slice := NewSlicer()
	slice.Add(xiaoMing)
	fmt.Println("current struct slice:", slice)
	slice.Add(xiaoLei)
	fmt.Println("current struct slice:", slice)
	slice.Add(xiaoLei)
	fmt.Println("current struct slice:", slice)
	slice.Add(xiaoFang)
	fmt.Println("current struct slice:", slice)
	slice.Remove(xiaoLei)
	fmt.Println("current struct slice:", slice)
	slice.Remove(xiaoLei)
	fmt.Println("current struct slice:", slice)
	slice.Remove(xiaoFang)
	fmt.Println("current struct slice:", slice)
	fmt.Println("struct slice end")
}

type Student struct {
	id   string
	name string
}

func (this Student) IsEqual(obj interface{}) bool {
	if student, ok := obj.(Student); ok {
		return this.GetId() == student.GetId()
	}
	panic("unexpected type")
}

func (this Student) GetId() string {
	return this.id
}
func intSliceExec() {
	fmt.Println("int slice start")
	slice := NewSlicer()
	slice.Add(1)
	fmt.Println("current int slice:", slice)
	slice.Add(2)
	fmt.Println("current int slice:", slice)
	slice.Add(2)
	fmt.Println("current int slice:", slice)
	slice.Add(3)
	fmt.Println("current int slice:", slice)
	slice.Remove(2)
	fmt.Println("current int slice:", slice)
	slice.Remove(2)
	fmt.Println("current int slice:", slice)
	slice.Remove(3)
	fmt.Println("current int slice:", slice)
	fmt.Println("int slice end")
}

func stringSliceExec() {
	fmt.Println("string slice start")
	slice := NewSlicer()
	slice.Add("hello")
	fmt.Println("current string slice:", slice)
	slice.Add("golang")
	fmt.Println("current string slice:", slice)
	slice.Add("golang")
	fmt.Println("current string slice:", slice)
	slice.Add("generic")
	fmt.Println("current string slice:", slice)
	slice.Remove("golang")
	fmt.Println("current string slice:", slice)
	slice.Remove("golang")
	fmt.Println("current string slice:", slice)
	slice.Remove("generic")
	fmt.Println("current string slice:", slice)
	fmt.Println("string slice end")
}
