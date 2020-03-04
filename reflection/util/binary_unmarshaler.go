package util

import (
	"encoding"
	"reflect"
)

func BinaryUnmarshaler(field reflect.Value) (b encoding.BinaryUnmarshaler) {
	InterfaceFrom(field, func(v interface{}, ok *bool) { b, *ok = v.(encoding.BinaryUnmarshaler) })
	return b
}
