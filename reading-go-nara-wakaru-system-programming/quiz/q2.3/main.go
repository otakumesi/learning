package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/test.json", handler)
	http.ListenAndServe(":8090", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Contet-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}

	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()

	gzipWriter.Header.Name = "sample.json"
	mWriter := io.MultiWriter(gzipWriter, os.Stdout)

	encoder := json.NewEncoder(mWriter)
	encoder.SetIndent("", "	")
	encoder.Encode(source)

	gzipWriter.Flush()
}
