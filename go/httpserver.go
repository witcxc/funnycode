package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {

	fmt.Printf("[%v]\n", req)
	buf := &bytes.Buffer{}
	nRead, err := io.Copy(buf, req.Body)
	if err != nil {
		fmt.Printf("\nerr[%v]\n", err)
	}
	fmt.Printf("\nbody[%v]\nlen[%v]\n", buf, nRead)
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		fmt.Printf("\nListenAndServe: %v\n", err)
	}
}
