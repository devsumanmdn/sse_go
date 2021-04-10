package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/lucas-clemente/quic-go/http3"
	"github.com/mitchellh/go-homedir"
)

func setupHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	wwwDir, _ := homedir.Expand("/Users/suman/go/sse_go/www")

	fmt.Println(wwwDir)

	mux.Handle("/", http.FileServer(http.Dir(wwwDir)))

	return mux
}

// Start a server that echos all data on the first stream opened by the client
func main() {

	certFile, _ := homedir.Expand("localhost.crt")
	keyFile, _ := homedir.Expand("localhost.key")

	handler := setupHandler()

	err := http3.ListenAndServe(":3000", certFile, keyFile, handler)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Println("Running on https://locahost:3000")
	}
}
