package util

import "unsafe"

// ByteToString ...
func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
