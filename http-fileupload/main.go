package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/upload", fileUploadHandler)
	fmt.Println("Server running on :8080 ")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fileUploadHandler(w http.ResponseWriter, r *http.Request){

	http.ServeFile(w,r,"templates/upload.html")
	// 10MB Limit
	r.ParseMultipartForm(10 << 20)
	
	//Retrieve file from form data
	file, handler, err := r.FormFile("myFile")
	if err!= nil {
		http.Error(w, "Error bad request", http.StatusBadRequest)
		return
	}
	defer file.Close()
	
	fmt.Fprintf(w, "Filename: %s\n", handler.Filename)
	fmt.Fprintf(w, "Filesize: %s\n", handler.Size)
	fmt.Fprintf(w, "MIME Header: %s\n", handler.Header)

	//Save file temp location
	dst, err := createFile(handler.Filename)
	if err != nil{
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	//Copy file to destination
	if _, err := dst.ReadFrom(file); err != nil {

		http.Error(w, "Error saving file", http.StatusInternalServerError)
	}
}

func createFile(filename string)(*os.File, error){
	//Create upladdir if not exist
	if _, err := os.Stat("uploads"); os.IsNotExist(err){
		os.Mkdir("uploads", 0755)
	}
	//build file path and createFile
	dst, err := os.Create(filepath.Join("uploads", filename))
	if err != nil {
		return nil, err
	}
	return dst, nil
}


