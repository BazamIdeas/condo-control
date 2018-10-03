package faces

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/astaxie/beego"
	"github.com/gofrs/uuid"
)

var (
	azureKey     = "xxxxx"
	azureBaseURL = "xxxxx"
)

var (
	baseImageURL    = beego.AppConfig.String("faces::baseUrl")
	rootDir, _      = filepath.Abs(beego.AppConfig.String("assets::jumps"))
	imageFolderPath = beego.AppConfig.String("assets::imageFolderPath")
	imageFolderDir  = rootDir + "/" + imageFolderPath
)

func init() {

	checkOrCreateImagesFolder(imageFolderDir)
}

func checkOrCreateImagesFolder(imageFolderDir string) (err error) {
	if _, err := os.Stat(imageFolderDir); os.IsNotExist(err) {
		os.MkdirAll(imageFolderDir, 644)
	}
	return
}

//CreateFaceFile create a image File
func CreateFaceFile(fh *multipart.FileHeader) (imageUUID string, mimeType string, err error) {

	file, err := fh.Open()

	if err != nil {
		return
	}

	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	mimeType = http.DetectContentType(fileBytes)

	if mimeType != "image/png" && mimeType != "image/jpeg" {
		err = errors.New("Bad mime-type")
		return
	}

	UUID, err := uuid.NewV4()

	if err != nil {
		return
	}

	fileName := UUID.String()

	err = ioutil.WriteFile(imageFolderDir+"/"+fileName, fileBytes, 0644)

	if err != nil {
		return
	}

	imageUUID = fileName

	return
}

func GetFaceFile(imageUUID string) (imageBytes []byte, err error) {

	imageURL := imageFolderDir + "/" + imageUUID

	imageBytes, err = ioutil.ReadFile(imageURL)

	return

}

//DeleteFaceFile ...
func DeleteFaceFile(imageUUID string) (err error) {

	err = os.Remove(imageUUID)

	return
}

//CreateFaceID ...
func CreateFaceID(imageUUID string) (faceID string, err error) {

	// Create the Http client
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	imageURL := baseImageURL + "/" + imageUUID

	requestBody := map[string]string{"url": imageURL}

	requestBodyBytes, err := json.Marshal(requestBody)

	if err != nil {
		return
	}

	bodyReader := bytes.NewReader(requestBodyBytes)

	req, err := http.NewRequest("POST", azureBaseURL+"/detect", bodyReader)
	if err != nil {
		return
	}

	req.Header.Add("Ocp-Apim-Subscription-Key", azureKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if res.StatusCode != 200 {
		err = errors.New("Face Request Failed")
		return
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)

	responseBodyJSON := []map[string]interface{}{}

	err = json.Unmarshal(bodyBytes, responseBodyJSON)

	if err != nil {
		return
	}

	if len(responseBodyJSON) == 0 {
		err = errors.New("No faces found on Request")
		return
	}

	if len(responseBodyJSON) > 1 {
		err = errors.New("Too many Faces, 1 Face allowed")
		return
	}

	faceID = responseBodyJSON[0]["faceId"].(string)

	return
}

//CompareFacesIDs ...
func CompareFacesIDs(oldFaceID string, newFaceID string) (ok bool, err error) {

	// Create the Http client
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	requestBody := map[string]string{"faceId1": oldFaceID, "faceId2": newFaceID}

	requestBodyBytes, err := json.Marshal(requestBody)

	if err != nil {
		return
	}

	bodyReader := bytes.NewReader(requestBodyBytes)

	req, err := http.NewRequest("POST", azureBaseURL+"/verify", bodyReader)

	if err != nil {
		return
	}

	req.Header.Add("Ocp-Apim-Subscription-Key", azureKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if res.StatusCode != 200 {
		err = errors.New("Face Request Failed")
		return
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)

	responseBodyJSON := map[string]interface{}{}

	err = json.Unmarshal(bodyBytes, responseBodyJSON)

	if err != nil {
		return
	}

	if len(responseBodyJSON) == 0 {
		err = errors.New("No faces found on Request")
		return
	}

	if len(responseBodyJSON) > 1 {
		err = errors.New("Too many Faces, 1 Face allowed")
		return
	}

	ok = responseBodyJSON["isIdentical"].(bool)

	return
}
