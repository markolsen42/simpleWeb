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
			var pre = "<!-- from "+ splits[i] +"-->\n"+ "<div style='border-style: dotted'>"
			var post = "</div>"+ "<!-- end " +splits[i] + "-->\n"
			out += pre + string(html) + post
		}
	}
	return out;
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	html, err := ioutil.ReadFile("content/main.html")
	check(err)
	var insertSomeMore = true
	var splits []string
	var inserted =string(html)
	for insertSomeMore {
		splits = strings.Split(inserted,"***")
		inserted = loadInserts(splits) 
		if !strings.Contains(inserted, "***"){
			insertSomeMore = false
		}
	}
	fmt.Fprintf(w, inserted, r.URL.Path[1:])
}
//ttd
//load splits recursively***DONE***
//add comments to show where the components come from***DONE***
//show insert blocks visually***DONE***
//make content editable
