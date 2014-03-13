package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
	"strconv"
)

var ServerStatus int = 200

func requestHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var content string
	var err error

	reqType := r.URL.Path[1:]
	content, err = getContent(reqType)
	// Use 500 for errors
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Echo Headers
	for k, v := range r.Header {
		w.Header().Add("X-" + k, v[0])
	}

	endTime := time.Now()
	w.Header().Add("ReadTime", endTime.Sub(startTime).String())
	fmt.Fprint(w, content)
}

func statusRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(ServerStatus)
}

func failureRequestHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if val,ok := values["status"]; ok {
		v, err := strconv.ParseInt(val[0], 10, 0)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ServerStatus = int(v)
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/status", statusRequestHandler)
	http.HandleFunc("/failure", failureRequestHandler)
	http.HandleFunc("/", requestHandler)
	fmt.Println("Running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}