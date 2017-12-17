package flatten

import (
	"reflect"
)

func flatten(val reflect.Value) []interface{} {
	switch val.Kind() {
	case reflect.Slice:
		res := make([]interface{}, 0)
		for i := 0; i < val.Len(); i++ {
			flat := flatten(val.Index(i).Elem())
			for _, item := range flat {
				res = append(res, item)
			}
		}
		return res
	case reflect.Int:
		intVal := int(val.Int())
		return []interface{}{intVal}
	default:
		return []interface{}{}
	}
}

func Flatten(val interface{}) []interface{} {
	return flatten(reflect.ValueOf(val))
}
