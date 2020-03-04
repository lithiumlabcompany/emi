package reflection

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"

	"emi/reflection/setter"
	"emi/reflection/util"
)

const TagName = "emi"

// ErrInvalidSpecification indicates that a specification is of the wrong type.
var ErrInvalidSpecification = errors.New("specification must be a struct pointer")

var gatherRegexp = regexp.MustCompile("([^A-Z]+|[A-Z]+[^A-Z]+|[A-Z]+)")
var acronymRegexp = regexp.MustCompile("([A-Z]+)([A-Z][^A-Z]+)")

// A ParseError occurs when an environment variable cannot be converted to
// the type required by a struct field during assignment.
type ParseError struct {
	FieldName string
	TypeName  string
	KeyName   string
	Value     string
	Err       error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("emi.Process: assigning %[1]s to %[2]s: converting '%[3]s' to type %[4]s. details: %[5]s", e.KeyName, e.FieldName, e.Value, e.TypeName, e.Err)
}

// GatherInfo gathers information about the specified struct
func GatherInfo(namespace string, spec interface{}) ([]VarInfo, error) {
	s := reflect.ValueOf(spec)

	if s.Kind() != reflect.Ptr {
		return nil, ErrInvalidSpecification
	}
	s = s.Elem()
	if s.Kind() != reflect.Struct {
		return nil, ErrInvalidSpecification
	}
	typeOfSpec := s.Type()

	// over allocate an info array, we will extend if needed later
	infos := make([]VarInfo, 0, s.NumField())
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		fieldType := typeOfSpec.Field(i)
		if !field.CanSet() || util.IsTrue(fieldType.Tag.Get("ignored")) {
			continue
		}

		for field.Kind() == reflect.Ptr {
			if field.IsNil() {
				if field.Type().Elem().Kind() != reflect.Struct {
					// nil pointer to a non-struct: leave it alone
					break
				}
				// nil pointer to struct: create a zero instance
				field.Set(reflect.New(field.Type().Elem()))
			}
			field = field.Elem()
		}

		// Capture information about the config variable
		info := VarInfo{
			Alt:   fieldType.Tag.Get(TagName),
			Tags:  fieldType.Tag,
			Name:  fieldType.Name,
			Field: field,
		}

		// Default to the field name as the env var name
		info.Key = info.Name

		if info.Alt != "" {
			info.Key = info.Alt
		}
		if namespace != "" {
			info.Key = fmt.Sprintf("%s.%s", namespace, info.Key)
		}

		infos = append(infos, info)

		if field.Kind() == reflect.Struct {
			if setter.From(field) == nil && util.TextUnmarshaler(field) == nil && util.BinaryUnmarshaler(field) == nil {
				innerNamespace := namespace
				if !fieldType.Anonymous {
					innerNamespace = info.Key
				}

				embeddedPtr := field.Addr().Interface()
				embeddedInfos, err := GatherInfo(innerNamespace, embeddedPtr)
				if err != nil {
					return nil, err
				}
				infos = append(infos[:len(infos)-1], embeddedInfos...)

				continue
			}
		}
	}

	return infos, nil
}
