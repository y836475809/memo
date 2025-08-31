package sv

import (
	"test/p1/sv"
	"test/p2/sv"
	"testing"
)

func TestCalc(t *testing.T) {
	sv.SetupRouter()
	sv.SetupRouter2()
	SetupRouter3()
	t.Error("test")
}
