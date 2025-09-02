package sample6_test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/h2non/gock"
)

type ResponseData struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func TestFoo(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("http://testdummyserver.com").
		Get("/bar").
		Reply(200).
		JSON(ResponseData{Message: "foo", Status: "Status"})

	r, err := http.Get("http://testdummyserver.com/bar")
	if err != nil {
		t.Error(err.Error())
	}
	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var data ResponseData
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		t.Error(err.Error())
	}
	println("data=", string(bodyBytes))
	if data.Message != "foo" || data.Status != "Status" {
		t.Error("data")
	}
}
