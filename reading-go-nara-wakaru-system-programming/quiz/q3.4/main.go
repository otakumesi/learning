package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	file, err := zipWriter.Create("hello.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(file, strings.NewReader("Hello, ZipWorld!!!"))
}
