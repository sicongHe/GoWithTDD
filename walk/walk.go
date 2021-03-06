package walk

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	//val := getValue(x)
	//if val.Kind() == reflect.String {
	//	fn(val.String())
	//	return
	//}
	//if val.Kind() == reflect.Slice {
	//	for i := 0; i < val.Len(); i++ {
	//		Walk(val.Index(i).Interface(), fn)
	//	}
	//	return
	//}
	//for i := 0; i < val.NumField(); i++ {
	//	field := val.Field(i)
	//	switch field.Kind() {
	//	case reflect.String:
	//		fn(field.String())
	//	case reflect.Struct:
	//		Walk(field.Interface(), fn)
	//	}
	//}
	val := getValue(x)
	numberOfValues := 0
	var getField func(int) reflect.Value
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			Walk(val.MapIndex(key).Interface(), fn)
		}
	}
	for i := 0; i < numberOfValues; i++ {
		Walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		return val.Elem()
	}
	return val
}
