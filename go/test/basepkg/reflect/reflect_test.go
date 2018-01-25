package basePkgT

import "reflect"

type myJob struct {
	t      reflect.Type
	args   []reflect.Value
	result []reflect.Value
}

func (f myJob) Run() {
	fn := reflect.MakeFunc(f.t, func(args []reflect.Value) (results []reflect.Value) {
		return reflect.New(f.t).Call(args)
	})
	f.result = fn.Call(f.args)
}
