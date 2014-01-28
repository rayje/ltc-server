package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var content string
	var err error

	reqType := r.URL.Path[1:]

	if reqType != "status" {
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
	}

	fmt.Fprint(w, content)
}

func main() {
	http.HandleFunc("/", requestHandler)
	fmt.Println("Running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}