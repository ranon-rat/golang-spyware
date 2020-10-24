package main

import (
	"bytes"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/kbinani/screenshot"
)

func sendI(w http.ResponseWriter, r *http.Request) {

	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, _ := screenshot.CaptureRect(bounds)
		buffer := new(bytes.Buffer)
		if err := png.Encode(buffer, img); err != nil {
			log.Println("unable to encode image.")
		}
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
		if _, err := w.Write(buffer.Bytes()); err != nil {
			log.Println("unable to write image.")
		}

	}
}

func main() {
	http.HandleFunc("/image", sendI)
	http.Handle("/", http.FileServer(http.Dir("./page")))

	http.ListenAndServe(":8090", nil)
}
