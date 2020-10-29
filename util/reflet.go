package util

import "reflect"

func ReflectValueToString(value reflect.Value) string {
	return value.String()
}
