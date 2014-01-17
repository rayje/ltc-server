package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	reqType := r.URL.Path[1:]

	startTime := time.Now()
	content, err := getContent(reqType)
	endTime := time.Now()

	// Use 500 for errors
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("ReadTime", endTime.Sub(startTime).String())

	// Echo Headers
	for k, v := range r.Header {
		w.Header().Add("X-" + k, v[0])
	}

	fmt.Fprint(w, content)
}

func main() {
	http.HandleFunc("/", requestHandler)
	fmt.Println("Running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}