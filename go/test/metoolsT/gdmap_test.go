package metoolsT

import (
	"testing"

	metools "github.com/ipiao/metools/thirdpartyAPI"
)

func TestGDMap(t *testing.T) {
	//res, err := metools.GdMapAddressToCoordHSON("杭州市拱墅区天堂E谷")
	res, err := metools.GdMapAddressToCoordStruct("淮安经济技术开发区南马厂乡开祥路5号")
	t.Log(res, err)
}
