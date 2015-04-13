package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestCallEcho(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Compile(w, r)
	}))
	defer ts.Close()

	//Creating Request
	values := make(url.Values)
	values.Set("doc", "{\"Type\":\"echo\", \"Code\": \"abc\", \"Output\": \"\"}")
	r, _ := http.PostForm(ts.URL, values)
	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	if string(body) != "abc" {
		t.Errorf("got %s want abc", body)
	}
}

// handling a POST connection with a variable named doc
func Compile(w http.ResponseWriter, req *http.Request) {
	var code map[string]string
	doc := req.FormValue("doc")
	fmt.Println("doc ", doc)
	err := json.Unmarshal([]byte(doc), &code)
	if err != nil {
		log.Printf("json.Unmarshal err = %s string = %s", err, doc)
		return
	}

	fmt.Println("code ", code["Type"])
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(code["Code"]))
}
