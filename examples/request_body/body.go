package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	bs, _ := ioutil.ReadAll(r.Body)
	var str strings.Builder

	str.Write(bs)
	str.WriteString("\n")
	fmt.Fprint(w, str.String())
	// reuse the request body with ioutil.NopCloser
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bs))
	bs, _ = ioutil.ReadAll(r.Body)
	fmt.Fprint(w, "another ioutil.ReadAll\n")
	fmt.Fprint(w, string(bs))
	r.Body.Close()
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
