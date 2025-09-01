package sv

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	sv1 "test/p1/sv"
	sv2 "test/p2/sv"
	"testing"
)

func TestRequest(t *testing.T) {
	r1 := sv1.SetupRouter1()
	r2 := sv2.SetupRouter2()
	r3 := SetupRouter3()

	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("GET", "/ping1", nil)
	r1.ServeHTTP(w1, req1)
	fmt.Printf("%v\n", w1.Code)
	fmt.Printf("%v\n", w1.Body.String())

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/ping2", nil)
	r2.ServeHTTP(w2, req2)
	fmt.Printf("%v\n", w2.Code)
	fmt.Printf("%v\n", w2.Body.String())

	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("GET", "/ping3", nil)
	r3.ServeHTTP(w3, req3)
	fmt.Printf("%v\n", w3.Code)
	fmt.Printf("%v\n", w3.Body.String())

	t.Error("test")
}
