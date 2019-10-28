package service

import (
	"fmt"
	"moges/domain/model"
	"moges/server/handlers/validators"
	"os"
	"regexp"
	"time"
)

func StorePhoto(request validators.PhotoRequest) error {

	createdTime := time.Now().Unix()
	fileName := fmt.Sprintf("%d%s", createdTime, request.Name)
	space := regexp.MustCompile(`\s+`)
	fileName = space.ReplaceAllString(fileName, "_")
	path := "/tmp/" + fileName;

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = f.Write(request.Data)
	if err != nil {
		return err
	}

	photo := model.Photo{
		Name:  fileName,
		Path:  path,
		Size:  request.Size,
	}
	err = photo.Save()

	if err != nil {
		_ = os.Remove(path)
	}
	defer f.Close()

	return err
}