package util

import (
	"encoding"
	"reflect"
)

func TextUnmarshaler(field reflect.Value) (t encoding.TextUnmarshaler) {
	InterfaceFrom(field, func(v interface{}, ok *bool) { t, *ok = v.(encoding.TextUnmarshaler) })
	return t
}
