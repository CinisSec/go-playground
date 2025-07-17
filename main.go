package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static"))) // Serve your HTML page
	http.HandleFunc("/cases/", uploadHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // Max 10 MB files
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	caseID := r.FormValue("case")
	if caseID == "" {
		http.Error(w, "Missing case ID", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	uploadDir := filepath.Join("cases", caseID, "uploads")

	os.MkdirAll(uploadDir, os.ModePerm)

	for _, fileHeader := range files {
		src, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Could not open file", http.StatusInternalServerError)
			return
		}
		defer src.Close()

		dstPath := filepath.Join(uploadDir, filepath.Base(fileHeader.Filename))

		dst, err := os.Create(dstPath)
		if err != nil {
			http.Error(w, "Could not create file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		if err != nil {
			http.Error(w, "Could not save file", http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprintf(w, "Uploaded %d file(s) to case: %s", len(files), caseID)
}
