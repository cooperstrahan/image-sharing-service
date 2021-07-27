package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/cooperstrahan/image-sharing-service/api/pkg/models"
)

var NewImage models.Image

func UploadImage(w http.ResponseWriter, r *http.Request) {
	indexHandler(w, r)

	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	imageTitle := r.FormValue("title")
	imageTags := r.FormValue("tags")
	imageDescription := r.FormValue("description")
	imageFileName := handler.Filename

	// fmt.Println(imageTitle)
	// fmt.Println(imageTags)
	// fmt.Println(imageDescription)
	// fmt.Println(imageFileName)

	image := models.Image{
		FileName:    imageFileName,
		Title:       imageTitle,
		Tags:        imageTags,
		Description: imageDescription,
	}

	image.UploadImage()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	// tempFile, err := ioutil.TempFile("images", handler.Filename)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	newFile := fmt.Sprintf("images/%s", handler.Filename)

	ioutil.WriteFile(newFile, fileBytes, 0777)
	// write this byte array to our temporary file
	// tempFile.Write(fileBytes)
	// return that we have successfully uploaded our

	w.WriteHeader(http.StatusOK)

}

func GetImages(w http.ResponseWriter, r *http.Request) {
	indexHandler(w, r)
	allImages := models.GetAllImages()
	res, _ := json.Marshal(allImages)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetImageById(w http.ResponseWriter, r *http.Request) {
	indexHandler(w, r)
	vars := mux.Vars(r)
	imageId := vars["id"]
	ID, err := strconv.ParseInt(imageId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	imageDetails, _ := models.GetImageById(ID)
	fmt.Printf(imageDetails.Title)
	res, _ := json.Marshal(imageDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	indexHandler(w, r)
	fmt.Println("Delete Image hit")
	vars := mux.Vars(r)
	imageId := vars["id"]
	ID, err := strconv.ParseInt(imageId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	models.DeleteImage(ID)
	// res, _ := json.Marshal(image)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// w.Write(res)
}

func IgnoreOptions(w http.ResponseWriter, r *http.Request) {
	indexHandler(w, r)
}

func enableCors(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	enableCors(&w, req)
	if req.Method == "OPTIONS" {
		fmt.Println("preflight")
		return
	}
}
