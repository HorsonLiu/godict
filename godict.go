package godict

import(
	"encoding/json"
	"reflect"
	"fmt"
)

type Dict map[string]interface{}

func mapToString(m map[string]interface{}) string {
	ret := ""

	for key, value := range m {
		if ret != "" {
			ret += ","
		} 

		ret += fmt.Sprintf("%s:%s", Str(key), Str(value))
	}

	return fmt.Sprintf("{%s}", ret)
}

func (this Dict)String() string {
	return mapToString(this)
}

func Str(obj interface{}) string {
	value := reflect.ValueOf(obj)
	kind := value.Kind()
	// fmt.Println("v=", value, "kind is ", kind)
	switch kind {
	default:
		return obj.(Dict).String()
	case reflect.Bool:
		if obj.(bool) {
			return "true"
		} else {
			return "false"
		}
	case reflect.Int:
		return fmt.Sprintf("%d", obj.(int))
	case reflect.Float32:
		return fmt.Sprintf("%f", obj.(float32))
	case reflect.Float64:
		return fmt.Sprintf("%f", obj.(float64))
	case reflect.String:
		return fmt.Sprintf("\"%s\"", obj.(string))
	case reflect.Slice:
		ret := ""
		for i := 0; i < value.Len(); i++ {
			if ret != "" {
				ret += ","
			}
			ret += Str(value.Index(i).Interface())
        }
		return fmt.Sprintf("[%s]", ret)
	case reflect.Array:
		ret := ""
		for i := 0; i < value.Len(); i++ {
			if ret != "" {
				ret += ","
			}
			ret += Str(value.Index(i).Interface())
        }
		return fmt.Sprintf("[%s]", ret)
	case reflect.Struct:
		// fmt.Println(obj)
		// fmt.Println(reflect.TypeOf(obj))
		ret, _ := json.Marshal(obj)
		return string(ret)
	case reflect.Ptr:
		payload := reflect.Indirect(value)
		// fmt.Println("palyload =", payload)
		return Str(payload.Interface())
	case reflect.Map:
		ret := ""
		keys := value.MapKeys()
		for _,key := range keys {
			if ret != "" {
				ret += ","
			}
			val := value.MapIndex(key)
			ret += fmt.Sprintf("%s:%s", Str(key.Interface()), Str(val.Interface()))
		}
		return fmt.Sprintf("{%s}", ret)
	case reflect.Int16:
		return fmt.Sprint(value.Interface().(int16))
	case reflect.Int32:
		return fmt.Sprint(value.Interface().(int32))
	case reflect.Int64:
		return fmt.Sprint(value.Interface().(int64))
	case reflect.Int8:
		return fmt.Sprint(value.Interface().(int8))
	case reflect.Uint64:
		return fmt.Sprint(value.Interface().(uint64))
	case reflect.Uint8:
		return fmt.Sprint(value.Interface().(uint8))
	case reflect.Uint16:
		return fmt.Sprint(value.Interface().(uint16))
	case reflect.Uint32:
		return fmt.Sprint(value.Interface().(uint32))
	case reflect.Uint:
		return fmt.Sprint(value.Interface().(uint))
	}
}
