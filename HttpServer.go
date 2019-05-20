package main

/**
Http服务器
*/
import (
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":12345", nil)
}
