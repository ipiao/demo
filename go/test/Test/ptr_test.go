package Test

import (
	"fmt"
	"testing"
)

type S struct {
	A int
}

func TestPtr(t *testing.T) {
	var s = new(S)
	// copy(s)
	var ss = &S{
		A: 1,
	}
	 s = ss
	fmt.Println(*s)
}

//func copy(s *S) {
//	var ss = &S{
//		A: 1,
//	}
//	s = ss
//	//s.A=1
//}
