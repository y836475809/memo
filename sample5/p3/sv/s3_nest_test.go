package sv

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	sv2 "test/p2/sv"
	"testing"
)

func TestNest(t *testing.T) {
	r2 := sv2.SetupRouter2()
	r3 := SetupRouter3()

	srv := &http.Server{
		Addr:    "localhost:8088",
		Handler: r2.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	defer srv.Shutdown(context.Background())

	// ping3-2 -> p2 /ping2
	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("GET", "/ping3-2", nil)
	r3.ServeHTTP(w3, req3)
	fmt.Printf("%v\n", w3.Code)
	fmt.Printf("%v\n", w3.Body.String())
	if w3.Body.String() != "pong2_pong3" {
		t.Error("not pong2_pong3")
	}
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Println("Server Shutdown:", err)
	// }
	fmt.Printf("%v\n", "test end")
}
