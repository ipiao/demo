package metoolsT

import (
	"testing"

	metools "github.com/ipiao/metools/thirdpartyAPI"
)

func TestAlimobile(t *testing.T) {
	// res, err := metools.AliMobileBelongToJSON("13777462204")
	// t.Log(res, err)
	res, err := metools.AliMobileBelongToStruct("13777462204")
	t.Log(res, err)
}

func TestAliZNWD(t *testing.T) {
	// res, err := metools.AliJiSuZNWDToJSON("今天天气如何")
	// t.Log(res, err)
	res, err := metools.AliJiSuZNWDToStruct("今天天气如何")
	t.Log(res, err)
}

func TestOcrIDCard(t *testing.T) {
	// res, err := metools.TransPicToOcrIDCardInput("../resources/images/idcard_face.jpg",
	// 	"face")
	// t.Log(res, err)
	res, err := metools.TransPicToOcrIDCardInputStr("../resources/images/idcard_face4.jepg",
		"face")
	t.Log(res, err)
	str, err := metools.AliOcrIDCardToJSON(res)
	t.Log(str, err)
	// t.Log(res, err)
}
