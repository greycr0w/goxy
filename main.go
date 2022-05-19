package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/greycr0w/goxy/headers"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("2/1/2006 15:04:05") + " [GOXY]:" + string(bytes))
}

func main() {
	address := "127.0.0.1"
	port := 9090
	fullAddress := fmt.Sprintf("%s:%d", address, port)
	fmt.Printf("Goxy listening on address: %s\n", fullAddress)

	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		httpMethod := fmt.Sprintf("%s %s%s\n", r.Method, r.Host, r.URL.String())
		body := string(bodyBytes)
		headers := headers.ParseHeaders(r.Header)
		log.Printf("\n\n%s\n%s\n\n%s\n\n", httpMethod, headers, body)

		//Do not return anything for now
		// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(fullAddress, nil))
	fmt.Println("hello")
}
