package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8090", nil)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	html, err := ioutil.ReadFile("test.html")
	check(err)
	fmt.Fprintf(w, string(html), r.URL.Path[1:])
}
