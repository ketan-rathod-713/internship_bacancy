package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

			// map[string][]string of the QueryString parameters.
			log.Println("QUERY PARAMS GOT", r.URL.Query())

			log.Println(r.URL.Query().Get("name"))
			if r.URL.Query().Get("name") == "ketan,aman" {
				log.Println("Name query is ketan")
			}

			w.Write([]byte("Hello World!"))
		}

		if r.Method == http.MethodPost {
			// for form data which is urlencoded
			r.ParseForm()
			log.Println(r.Form) // This will get whole information from form to query parameters ha ha great and hence we can use it.
			log.Println(r.URL.Query())

			w.Write([]byte(fmt.Sprintf("Hello I am server responding to POST request.")))
		}
	})

	// Response body field is streamed on demand as it is read.

	log.Fatal(http.ListenAndServe(":8080", nil))

}
