package web

//func Populate(dest interface{}) error {
//	value := reflect.ValueOf(dest)
//	if value.Kind() != reflect.Ptr {
//		return errors.New("must pass a point not a value")
//	}
//	if value.IsNil() {
//		return errors.New("nil point passed to populate")
//	}
//	mutable := value.Elem()
//	for i := 0; i < mutable.NumField(); i++ {
//		field := mutable.Field(i)
//		logs.Debugf("%v", field.Type())
//		switch field.Type().String() {
//		case "string":
//			field.SetString("test")
//		case "zero.String":
//			field.Set()
//		}
//	}
//	//	typ := reflect.TypeOf(value)
//	//	for i := 0; i < typ.NumField(); i++ {
//	//		logs.Debugf("%v", typ.Field(i))
//	//	}
//	return nil
//}
