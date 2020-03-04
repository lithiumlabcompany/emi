package setter

import (
	"reflect"

	"emi/reflection/util"
)

func From(field reflect.Value) (s Setter) {
	util.InterfaceFrom(field, func(v interface{}, ok *bool) { s, *ok = v.(Setter) })
	return s
}
