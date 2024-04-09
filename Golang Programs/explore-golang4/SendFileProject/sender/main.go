package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"net"
	"path/filepath"

	"fileshare/models"
)

func main() {
	var folderDirPath string

	flag.StringVar(&folderDirPath, "dir", "/home/bacancy/Desktop/Bacancy/internship_bacancy/Golang Programs/explore-golang4/SendFileProject/sender", "Specify folder path to send")
	flag.Parse()

	fmt.Println("Folder Directory Path:", folderDirPath)

	dirName := filepath.Dir(folderDirPath)
	fmt.Println("Directory Name:", dirName)

	var folder models.Folder = models.Folder{
		Name: filepath.Base(folderDirPath),
	}

	var files []models.File

	filepath.Walk(folderDirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file := models.File{Name: info.Name(), Size: int(info.Size()), ContentType: info.Mode().Type().String()}
			files = append(files, file)
		}
		return nil
	})

	folder.Files = files

	fmt.Println("Folder Metadata:", folder)

	// Start TCP server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		go handleClient(conn, folder)
	}
}

func handleClient(conn net.Conn, folder models.Folder) {
	defer conn.Close()

	// Encode folder metadata and files
	data, err := json.Marshal(folder)
	if err != nil {
		fmt.Println("Error encoding data:", err.Error())
		return
	}

	// Send data to client
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error sending data:", err.Error())
		return
	}
}
