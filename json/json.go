package json

import (
	"encoding/json"
	"errors"
	"github.com/withf/art/slices"
	"reflect"
)

// StructMarshal 用于输出结构体的JSON表示
// @param obj interface{} 结构体对象
// @param names []string 字符串切片，表明想要输出的字段
// @return []byte, error
func StructMarshal(obj interface{}, names ...string) ([]byte, error) {
	return structMarshal(true, obj, names...)
}

// StructMarshalExclude 用于输出结构体的JSON表示
// @param obj interface{} 结构体对象
// @param names []string 字符串切片，表明想不要输出的字段
// @return []byte, error
func StructMarshalExclude(obj interface{}, names ...string) ([]byte, error) {
	return structMarshal(false, obj, names...)
}

func structMarshal(include bool, obj interface{}, names ...string) ([]byte, error) {
	v, f := isStruct(obj)
	if !f {
		return nil, errors.New("The first param must be a Struct.")
	}

	if len(names) == 0 {
		return json.Marshal(obj)
	}

	t := reflect.TypeOf(obj)
	m := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		name := field.Name
		if include == slices.Contains(names, name) {
			tag := field.Tag.Get("json")
			if tag == "" {
				tag = name
			}
			m[tag] = v.Field(i).Interface()
		}
	}
	return json.Marshal(m)
}

func isStruct(obj interface{}) (reflect.Value, bool) {
	v := reflect.ValueOf(obj)
	return v, v.Kind() == reflect.Struct
}
