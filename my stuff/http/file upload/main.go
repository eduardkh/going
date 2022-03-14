package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// https://www.youtube.com/watch?v=7KzRfBsMdQQ&list=PLDZ_9qD1hkzOQdLHOPHtDcxoDSr0nno9G&index=14
// https://tutorialedge.net/golang/go-file-upload-tutorial/
var templates = template.Must(template.ParseGlob("templates/*.gohtml"))

func main() {
	// routes
	http.HandleFunc("/", index)
	http.HandleFunc("/fileupload", fileupload)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

// handlers
func index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.gohtml", nil)
}
func fileupload(w http.ResponseWriter, r *http.Request) {
	// GET give back the upload form
	if r.Method == "GET" {
		templates.ExecuteTemplate(w, "fileupload.gohtml", nil)
		return
	}
	// POST process the uploaded file
	log.Println("File Upload Endpoint Hit")
	r.ParseMultipartForm(10)
	file, fileHeader, err := r.FormFile("uploadfile")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	// log.Printf("Uploaded File: %+v\n", fileHeader.Filename)
	// log.Printf("File Size: %+v\n", fileHeader.Size)
	// log.Printf("MIME Header: %+v\n", fileHeader.Header)
	var osFile *os.File
	// GitHub copilot version (os.OpenFile)
	// osFile, err = os.OpenFile("./uploads/"+fileHeader.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// defer osFile.Close()
	// _, err = io.Copy(osFile, file)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// ioutil.TempFile version
	osFile, err = ioutil.TempFile("./uploads/", "*-"+fileHeader.Filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer osFile.Close()
	// get the new file name
	log.Println(osFile.Name())
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = osFile.Write(fileBytes)
	if err != nil {
		log.Println(err)
		return
	}
	http.Redirect(w, r, "/fileupload", http.StatusSeeOther)
	templates.ExecuteTemplate(w, "fileupload.gohtml", nil)
}
