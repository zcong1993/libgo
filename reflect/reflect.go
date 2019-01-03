package reflecthelper

import "reflect"

// IndirectType get the true type of a reflect Type
func IndirectType(in reflect.Type) reflect.Type {
	for in.Kind() == reflect.Ptr || in.Kind() == reflect.Slice {
		in = in.Elem()
	}
	return in
}

// IndirectValue get the true value of a reflect Value
func IndirectValue(in reflect.Value) reflect.Value {
	for in.Kind() == reflect.Ptr {
		in = in.Elem()
	}
	return in
}

// Slice2Map return a map[interface{}]interface{} which the key is the value of given filedName
func Slice2Map(sl interface{}, fieldName string) map[interface{}]interface{} {
	if reflect.TypeOf(sl).Kind() != reflect.Slice {
		panic("only support slice")
	}

	v := reflect.ValueOf(sl)

	if IndirectType(v.Type()).Kind() != reflect.Struct {
		panic("only support struct slice")
	}

	l := reflect.ValueOf(sl).Len()

	res := make(map[interface{}]interface{}, l)

	for i := 0; i < l; i++ {
		vv := IndirectValue(v.Index(i))
		value := vv.FieldByName(fieldName)
		if !value.IsValid() {
			panic("field not valid")
		}

		finalValue := IndirectValue(value).Interface()
		res[finalValue] = vv.Interface()
	}

	return res
}
