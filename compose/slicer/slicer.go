package slicer

import (
	"errors"
	"fmt"
)

type Slice []interface{}

func NewSlicer() Slice {
	return make(Slice, 0)
}

//通过这个比对，支持Comparable和golang基础类型
func isEqual(a, b interface{}) bool {
	if comparable, ok := a.(Comparable); ok {
		return comparable.IsEqual(b)
	} else {
		return a == b
	}
}

func (this *Slice) Add(elem interface{}) error {
	for _, v := range *this {
		if isEqual(v, elem) {
			fmt.Printf("Slice:Add elem: %v already exist\n", elem)
			return errors.New("slicer element already exist")
		}
	}
	*this = append(*this, elem)
	fmt.Printf("Slice:Add elem: %v succ\n", elem)
	return nil
}

func (this *Slice) Remove(elem interface{}) error {
	found := false
	for i, v := range *this {
		if isEqual(v, elem) {
			if i == len(*this)-1 {
				*this = (*this)[:i]

			} else {
				*this = append((*this)[:i], (*this)[i+1:]...)
			}
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Slice:Remove elem: %v not exist\n", elem)
		return errors.New("slicer element not exist")
	}
	fmt.Printf("Slice:Remove elem: %v succ\n", elem)
	return nil
}
