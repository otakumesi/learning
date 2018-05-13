package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			req, err := http.ReadRequest(
				bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}

			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dump))

			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
