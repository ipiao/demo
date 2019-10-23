package basePkgT

import (
	"errors"
	"testing"
)




func TestF(t *testing.T){
	var err error
	var a  string

	if b,err:=function2();err!=nil{
		a,err = function3(b)
	}
	t.Log(err,a)
}

func function2()(string,error){
	return "function2",errors.New("error function2")
}

func function3(string)(string,error){
	return "function3",errors.New("error function3")
}