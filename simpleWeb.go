package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var componentDelimiter = "***"

func main() {
	http.HandleFunc("/", initialHandler)
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

func initialHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		handlePost(w, r)

	} else {
		helloServer(w, r)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	// should be starts with not contains TODO
	changeEndpoint := "/change"
	if strings.Contains(r.URL.Path, changeEndpoint) {
		// get the posted body as a map
		result := getBodyMap(r)
		// print the html element of the body
		fmt.Print(result["html"])
		// get the part of the URL after /change this will be the name of the file written
		afterChange := r.URL.Path[len(changeEndpoint)+1:]
		// turn the html element of the body a interface{} into a string then a byte array
		d1 := []byte(fmt.Sprintf("%v", result["html"]))
		//write the file
		ioutil.WriteFile("content/"+afterChange, d1, 0644)

		jsonResp := "{\"Hello\":\"World\"}"
		fmt.Fprintf(w, jsonResp)
	}
}

//get the bodyMap passed to the endpoint
func getBodyMap(r *http.Request) map[string]interface{} {
	var result map[string]interface{}
	json.NewDecoder(r.Body).Decode(&result)
	return result
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

func helloServer(w http.ResponseWriter, r *http.Request) {
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
// test with multiple main input files ***DONE*** main and main2
//make an endpoint that changes a file? POST /change/main.html***WORKING***
// make a test post endpoint ***DONE*** initial handler to handlePost
//set response type to json
// make the postwork on a certain path /change
