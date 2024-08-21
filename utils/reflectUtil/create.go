package reflectUtil

import "reflect"

// NewStructPtr return a pointer to struct, generic T is a struct pointer
func NewStructPtr[T any]() T {
	tType := reflect.TypeOf((*T)(nil)).Elem() // 获取 T 的类型
	if tType.Kind() != reflect.Ptr || tType.Elem().Kind() != reflect.Struct {
		panic("T must be a pointer to a struct")
	}
	newStructPtr := reflect.New(tType.Elem()).Interface().(T)
	return newStructPtr
}

// SetGenericStructPointerDefault 泛型函数，接收一个指向类型T的指针变量并将其设置为T的默认值
func SetGenericStructPointerDefault[T any](value *T) {
	// 使用反射获得类型T的值
	v := reflect.ValueOf(value).Elem()

	// 检查T是否是指针类型
	if v.Kind() == reflect.Ptr {
		// 获得指针指向的元素类型
		elemType := v.Type().Elem()
		// 创建指向元素类型的新指针
		newValue := reflect.New(elemType).Interface()
		// 将这个新指针赋值给变量
		v.Set(reflect.ValueOf(newValue))
	} else {
		// 如果不是指针类型，则设置为零值
		v.Set(reflect.Zero(v.Type()))
	}
}
