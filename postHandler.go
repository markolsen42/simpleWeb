package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	// should be starts with not contains TODO

	changeEndpoint := "/change"
	if strings.Contains(r.URL.Path, changeEndpoint) {

		//this is possibly how to get data from application/x-www-form-urlencoded data - eg form post submit
		r.ParseForm()
		fmt.Println(r.Form)

		// get the posted body as a map - this worked for a json post from postman
		result := getBodyMap(r)
		// print the html element of the body
		fmt.Print(result)
		fmt.Print(result["html"])
		// get the part of the URL after /change this will be the name of the file written
		afterChange := r.URL.Path[len(changeEndpoint)+1:]
		// turn the html element of the body a interface{} into a string then a byte array
		//d1 := []byte(fmt.Sprintf("%v", result["html"]))
		d1 := []byte(r.Form["html"][0])
		//write the file
		ioutil.WriteFile("content/"+afterChange, d1, 0644)

		jsonResp := `<!DOCTYPE html>
		<html>
		   <head>
			  <title>HTML Meta Tag</title>
			  <meta http-equiv = "refresh" content = "2; url = /main.html" />
		   </head>
		   <body>
			  <p>Working</p>
		   </body>
		</html>`
		fmt.Fprintf(w, jsonResp)
	}
}
