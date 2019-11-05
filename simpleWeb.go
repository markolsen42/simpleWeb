package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var componentDelimiter = "***"

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8090", nil)
}

func check(e error, path string) {
	if e != nil && !strings.Contains(path, "favi") {
		switch e {
		case os.ErrInvalid:
			fmt.Println("ErrInvalid " + path)
		case os.ErrPermission:
			fmt.Println("ErrPermission " + path)
		case os.ErrNotExist:
			fmt.Println("ErrNotExist " + path)
		default:
			panic(e)
		}
	}
}


func loadInserts(splits []string, addScaffolding bool) string {
	var out = ""
	for i := 0; i < len(splits); i++ {
		if i%2 == 0 {
			out += splits[i]
		} else {
			html, err := ioutil.ReadFile("content/" + splits[i] + ".html")
			check(err, splits[i])
			if addScaffolding {
				out += formatInsert(splits[i], string(html))
			} else {
				out += string(html)
			}
		}
	}
	return out
}

func formatInsert(insertToken string, insertHtml string) string {
	var pre = "<!-- from " + insertToken + "-->\n" + "<div style='border-style: dotted'>"
	var post = "</div>" + "<!-- end " + insertToken + "-->\n"
	return pre + insertHtml + post
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	var path = r.URL.Path
	if path == "/" {
		path = "/main.html"
	}
	html, err := ioutil.ReadFile("content" + path)
	check(err, path)
	var addScaffolding = true
	var insertSomeMore = true
	var splits []string
	var inserted = string(html)
	for insertSomeMore {
		splits = strings.Split(inserted, componentDelimiter)
		inserted = loadInserts(splits, addScaffolding)
		if !strings.Contains(inserted, componentDelimiter) {
			insertSomeMore = false
		}
	}
	fmt.Fprintf(w, inserted, r.URL.Path[1:])
}

//ttd
//load splits recursively***DONE***
//add comments to show where the components come from***DONE***
//show insert blocks visually***DONE***
//make content editable***
