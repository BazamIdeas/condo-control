package files

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/gofrs/uuid"
)

var (
	rootDir, _ = filepath.Abs(beego.AppConfig.String("assets::jumps"))
)

func checkOrCreateFolder(folderName string) (folderPath string, err error) {

	folderPath = rootDir + "/assets/" + folderName

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, 644)
		err = nil
	}

	return
}

func getFilePath(folderName string, fileName string) (filePath string, err error) {

	folderPath, err := checkOrCreateFolder(folderName)

	if err != nil {
		return
	}

	filePath = folderPath + "/" + fileName

	return

}

//CreateFile create a image File
func CreateFile(fh *multipart.FileHeader, groupName string) (fileUUID string, mimeType string, err error) {

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

	filePath, err := getFilePath(groupName, fileName)

	if err != nil {
		return
	}

	err = ioutil.WriteFile(filePath, fileBytes, 0644)

	if err != nil {
		return
	}

	fileUUID = fileName

	return
}

func GetFile(fileUUID string, groupName string) (fileBytes []byte, mimeType string, err error) {

	filePath, err := getFilePath(groupName, fileUUID)

	if err != nil {
		return
	}

	fileBytes, err = ioutil.ReadFile(filePath)

	if err != nil {
		return
	}

	mimeType = http.DetectContentType(fileBytes)

	return

}

//DeleteFaceFile ...
func DeleteFile(fileUUID string, groupName string) (err error) {

	err = os.Remove(fileUUID)

	return
}
