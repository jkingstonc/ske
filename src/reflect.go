package src

import "reflect"

// get the struct field value by index
func FieldIndex(index int, value reflect.Value) reflect.Value{
	return value.FieldByIndex([]int{index})
}

// get the struct field value by name
func Field(name string, value reflect.Value) reflect.Value{
	return value.FieldByName(name)
}

// return the literal value of 'val'
func Value(val interface{}) reflect.Value{
	return reflect.ValueOf(val)
}

// return the literal value pointed to by 'val'
func ValuePtr(val interface{}) reflect.Value{
	return reflect.ValueOf(val).Elem()
}