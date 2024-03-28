package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			err := r.ParseMultipartForm(1024 * 1024 * 20)

			if err != nil {
				http.Error(w, "error parsing multipart form", http.StatusBadRequest)
			}

			file, handler, err := r.FormFile("file")
			if err != nil {
				http.Error(w, "Error retrieving file", http.StatusBadRequest)
				return
			}

			res, err := http.Post("https://store1.gofile.io/contents/uploadfile", "mutlipart/form-data", file)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(res.Status)

			fmt.Println(handler.Filename, handler.Size)

			w.Write([]byte("done"))
		}
	})

	http.ListenAndServe(":8081", nil)
}
