package models

type Folder struct {
	Name  string `json:"name"`
	Files []File `json:"files"`
}

type File struct {
	Name        string `json:"name"`
	ContentType string `json:"contentType"`
	Size        int    `json:"size"`
}
