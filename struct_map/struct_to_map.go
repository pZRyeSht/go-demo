package struct_map

import (
	"errors"
	"reflect"
)

// StructSliceToMapSlice : struct切片转为map切片
func StructSliceToMapSlice(source interface{}) (mpList []map[string]interface{}, err error) {
	v := reflect.ValueOf(source)
	if v.Kind() != reflect.Slice {
		return nil, errors.New(" Unknown type, slice expected.")
	}
	l := v.Len()
	// 将 interface 转为 []interface{}
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	// 通过遍历，每次迭代将struct转为map
	for _, elem := range ret {
		toMap, _ := StructToMap(elem)
		mpList = append(mpList, toMap)
	}
	return mpList, err
}

func StructToMap(source interface{}) (map[string]interface{}, error) {
	objV := reflect.ValueOf(source)
	if objV.Kind() == reflect.Ptr {
		objV = objV.Elem()
	}
	// 非结构体返回错误提示
	if objV.Kind() != reflect.Struct {
		return nil, errors.New("ToMap only accepts struct or struct pointer")
	}
	objT := objV.Type()
	mp := make(map[string]interface{}, 0)
	for i := 0; i < objV.NumField(); i++ {
		switch objV.Field(i).Kind() {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
			mp[objT.Field(i).Name] = objV.Field(i).Int()
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
			mp[objT.Field(i).Name] = objV.Field(i).Uint()
		case reflect.Float32, reflect.Float64:
			mp[objT.Field(i).Name] = objV.Field(i).Float()
		case reflect.String:
			mp[objT.Field(i).Name] = objV.Field(i).String()
		case reflect.Bool:
			mp[objT.Field(i).Name] = objV.Field(i).Bool()
		default:
		}
	}
	return mp, nil
}
