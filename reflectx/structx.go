package reflectx

import (
	"reflect"
)

// TypeOfStruct 直接获取struct的type
func TypeOfStruct(i interface{}) reflect.Type {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

// ValueOfStruct 直接获取struct的value
func ValueOfStruct(i interface{}) reflect.Value {
	t := reflect.ValueOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

// StructName 直接获取struct的name
func StructName(i interface{}) string {
	return TypeOfStruct(i).Name()
}

// StructRange 循环遍历struct的属性，使用handler处理相对应的属性
func StructRange(i interface{}, handler func(t reflect.StructField, v reflect.Value) error) error {
	v := ValueOfStruct(i)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		var (
			fv = v.Field(i)
			ft = t.Field(i)
		)
		if fv.CanInterface() { //过滤不可访问的属性
			if err := handler(ft, fv); err != nil {
				return err
			}
		}
	}
	return nil
}

// Empty 判断某一属性是否为空值
func Empty(i interface{}) bool {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		v := reflect.ValueOf(i)
		return v.IsNil()
	}
	zero := reflect.New(t).Elem()
	return reflect.DeepEqual(i, zero.Interface())
}
