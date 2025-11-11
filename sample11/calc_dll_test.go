package main

import (
	"testing"
)

// テストコード内に import "C"するとエラー
// cgoはテストコード内では使えない様子
// なのでgo側でラップした関数を呼ぶ

func TestSum(t *testing.T) {
	tmp := 10
	println(tmp)
	sum := calcSum()
	if sum != 100 {
		t.Error("error")
	}
}
