package sv

import (
	"test/p1/sv"
	"testing"
)

func TestCalc(t *testing.T) {
	sv.SetupRouter()
	SetupRouter2()
	t.Error("test")
}
