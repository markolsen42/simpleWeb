package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func loadInserts(splits []string) string {
	var out = "";
	for i:=0;i< len(splits);i++{
		if i%2 == 0 {
			out += splits[i]
		} else {
			html, err := ioutil.ReadFile("content/" + splits[i] + ".html")
			check(err)
			var pre = "<div style=\"border-style:solid\""
			var post = "</div>"
			out += pre+string(html)+ post
		}
	}
	return out;
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	html, err := ioutil.ReadFile("content/main.html")
	check(err)
	var splits = strings.Split(string(html),"***")
	fmt.Println(splits)
	var inserted = loadInserts(splits)
	fmt.Println(inserted)
	fmt.Fprintf(w, inserted, r.URL.Path[1:])
}
