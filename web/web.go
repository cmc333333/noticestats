package web

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func Run() {
	http.HandleFunc("/", hello)
	host := os.Getenv("OPENSHIFT_GO_IP")
	port := os.Getenv("OPENSHIFT_GO_PORT")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "8080"
	}
	bind := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Testing: hello, world from %s", runtime.Version())
}
