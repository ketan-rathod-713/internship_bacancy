package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	// Open the file that you want to upload
	file, err := os.Open("index.html")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "https://store1.gofile.io/contents/uploadfile", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Create a new multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the file to the form
	part, err := writer.CreateFormFile("file", "index.html")
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

	// Add any additional form fields
	_ = writer.WriteField("folderId", "1422707f-3ff7-43d8-8365-fb4f67bbd18c")

	// Close the multipart form
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing writer:", err)
		return
	}

	// Set the content type header
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Set the authorization header
	req.Header.Set("Authorization", "guest3734854153")

	// Set the request body
	req.Body = ioutil.NopCloser(body)

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println(string(responseBody))
}
