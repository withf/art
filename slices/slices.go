package slices

import (
	"errors"
	"reflect"
)

// Contains 切片中是否包含指定数据
// @param s interface{} 切片类型
// @param v interface{} 查找的数据
// @return bool
func Contains(s interface{}, v interface{}) bool {
	sli, err := convertInterfaceToSlice(s)
	if err != nil {
		return false
	}
	for _, t := range sli {
		if t == v {
			return true
		}
	}
	return false
}

func isSlice(s interface{}) (reflect.Value, bool) {
	v := reflect.ValueOf(s)
	return v, v.Kind() == reflect.Slice
}

func convertInterfaceToSlice(s interface{}) ([]interface{}, error) {
	v, f := isSlice(s)
	if !f {
		return nil, errors.New("The first param is must be a Slice.")
	}
	l := v.Len()
	sli := make([]interface{}, l)
	for i := 0; i < l; i++ {
		sli[i] = v.Index(i).Interface()
	}
	return sli, nil
}
