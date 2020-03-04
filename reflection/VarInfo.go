package reflection

import (
	"reflect"
)

// VarInfo maintains information about the configuration variable
type VarInfo struct {
	Name  string
	Alt   string
	Key   string
	Tags  reflect.StructTag
	Field reflect.Value
}
