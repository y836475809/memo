package sv

import (
	"test/p1/sv"
	"testing"
)

func TestCalc(t *testing.T) {
	sv.SetupRouter1()
	SetupRouter2()
	t.Error("test")
}
