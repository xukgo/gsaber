package ruleUtil

import "reflect"

func InterfaceIsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	kind := vi.Kind()
	if kind == reflect.Ptr {
		return vi.IsNil()
	} else if kind == reflect.Map {
		return vi.IsNil()
	} else if kind == reflect.Slice {
		return vi.IsNil()
	}
	return false
}
