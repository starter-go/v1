package platforms

import "testing"

func TestInfoBuilder(t *testing.T) {

	ib := new(InfoBuilder)
	loader := new(innerCurrentPlatformInfoLoader)

	loader.OnLoad(ib)

	info := ib.Info()
	t.Log(info)
	t.Log(ib)

}
