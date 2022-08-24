package dashboard

import (
	"io"
	"log"
	"net/http"
)

func Start() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/hello", helloHandler)

	log.Println("Web dashboard is running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
