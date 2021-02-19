package util

import (
	"reflect"
	"unsafe"

)

// GetType ...
func GetType(value interface{}) string {
	return reflect.TypeOf(value).Name()
}

// ByteToString ...
func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
