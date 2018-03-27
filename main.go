package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"time"
)

const (
	maximumChunkSizeBytes = 256000
)

func main() {
	http.HandleFunc("/manifest.m3u8", manifestHandler)
	http.HandleFunc("/ts/", segmentHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func manifestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving manifest")
	dat, err := ioutil.ReadFile("assets/manifest.m3u8")
	if err != nil {
		log.Fatalf("Error reading manifest file, exiting: %v", err)
	}
	w.Header().Add("Content-Type", "application/x-mpegURL")
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Write(dat)
}

func segmentHandler(w http.ResponseWriter, r *http.Request) {
	segmentName := r.URL.Path
	log.Printf("Received request: %s\n", segmentName)

	dat, err := ioutil.ReadFile(fmt.Sprintf("assets%s", segmentName))
	if err != nil {
		log.Printf("Error reading data file, exiting: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "video/MP2T")
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	chunks := int(math.Ceil(float64(len(dat)) / float64(maximumChunkSizeBytes)))
	for index := 0; index < chunks; index++ {
		begin := index * maximumChunkSizeBytes
		end := ((index + 1) * maximumChunkSizeBytes)
		if end > len(dat) {
			end = len(dat)
		}
		w.Write(dat[begin:end])
		time.Sleep(time.Millisecond * 500)
	}
}
