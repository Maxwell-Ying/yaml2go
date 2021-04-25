package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"yaml2go/task"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		defaultHandler(w, r)
		return
	}
	bodyString := string(bodyBytes)

	if !strings.Contains(bodyString, "=") {
		defaultHandler(w, r)
		return
	}

	bodyString = strings.Split(bodyString, "=")[1]

	bodyString, err = url.QueryUnescape(bodyString)

	if err != nil {
		defaultHandler(w, r)
		return
	}

	if bodyString != "" {
		w.Write([]byte(task.Convert(bodyString, "\t")))
		return
	}
	defaultHandler(w, r)

}

func defaultHandler(w http.ResponseWriter, r *http.Request)  {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":9527", nil)
	//http.HandleFunc("/convert", indexHandler)
//	result := task.Convert(
//		`
//test:
//  test1:
//    test2: aaa
//    test3: bbb
//  test4:
//    test5: ccc
//  test6: ddd
//test0: eee
//`, "    ")
//	println(result)
}
