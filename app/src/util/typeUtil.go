package util

import "reflect"

// GetType ...
func GetType(value interface{}) string {
	return reflect.TypeOf(value).Name()
}
