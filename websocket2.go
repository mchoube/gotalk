package main

import (
	"compress/gzip"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

func dnldHandler(ws *websocket.Conn) {
	file, err := os.Open("beautiful.jpg")
	if err != nil {
		http.Error(w, "file beautiful.jpg does not exist.", http.StatusBadRequest)
		return
	}
	defer file.Close()

	gz := gzip.NewWriter(ws)
	defer gz.Close()

	io.Copy(gz, file)
}
// STOP OMIT

func main() {
	http.Handle("/dnld", websocket.Handler(dnldHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
