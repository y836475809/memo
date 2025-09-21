package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/h2non/gock"
)

func TestSimpleGet(t *testing.T) {
	defer gock.Off()

	gock.New("http://foo.com").
		Get("/bar").
		MatchParam("a", `^1{1}$`).
		MatchParam("b", `^2{1}$`).
		Reply(200).
		JSON(map[string]string{"ID": "0", "Name": "hoge"})

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/ping", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = request

	StartGet(c)
	// fmt.Print()
	act := w.Body.String()
	exp := `{"key":"{\"ID\":\"0\",\"Name\":\"hoge\"}\n"}`
	if act != exp {
		t.Errorf("act=%v, exp=%v", act, exp)
	}
}

func TestSimplePost(t *testing.T) {
	defer gock.Off()

	gock.New("http://foo.com").
		Post("/post").
		MatchParam("a", `^1{1}$`).
		MatchParam("b", `^(10){1}$`).
		MatchType("json").
		JSON(Sample{ID: 0, Name: "hoge"}).
		Reply(200).
		JSON(map[string]string{"Result": "OK"})

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/ping", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = request

	StartPost(c)
	act := w.Body.String()
	exp := `{"key":"{\"Result\":\"OK\"}\n"}`
	if act != exp {
		t.Errorf("act=%v, exp=%v", act, exp)
	}
}
