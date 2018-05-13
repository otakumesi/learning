package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	request, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	request.Write(conn)
	res, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
