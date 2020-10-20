package tranquility

import "reflect"

// Tranquility returns new function execute side effect.
func Tranquility(p interface{}, f interface{}, g interface{}) {
	newF := tranquilityT(f, g)
	reflect.ValueOf(p).Elem().Set(newF)
}

func tranquilityT(f interface{}, g interface{}) reflect.Value {
	fType := reflect.TypeOf(f)
	return reflect.MakeFunc(fType, func(args []reflect.Value) []reflect.Value {
		reflect.ValueOf(g).Call(args)
		return reflect.ValueOf(f).Call(args)
	})
}
