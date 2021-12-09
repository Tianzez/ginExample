package utils

import (
	"reflect"
)

// Gets element of the pointer obj points to.
func elem(obj reflect.Value) reflect.Value {
	for obj.Kind() == reflect.Ptr {
		if obj.Elem().Kind() == reflect.Invalid {
			obj = reflect.New(obj.Type().Elem())
		} else {
			obj = obj.Elem()
		}
	}
	return obj
}

// Returns the type of interface obj or that the pointer obj points to.
func TypeOf(obj interface{}) reflect.Type {
	t := reflect.TypeOf(obj)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

// Returns the value of interface obj or that the pointer obj points to.
func ValueOf(obj interface{}) reflect.Value {
	v, ok := obj.(reflect.Value)
	if ok {
		return v
	} else {
		v = reflect.ValueOf(obj)
	}
	return elem(v)
}

func StructName(stc interface{}) string {
	return TypeOf(stc).Name()
}
