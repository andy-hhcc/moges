package validators

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type PhotoRequest struct {
	Data []byte
	Name string
	Size int
	Path string
}
func NewPhotoValidator() PhotoRequest {
	return PhotoRequest{}
}

func (validator *PhotoRequest) Bind(c *gin.Context) error {
	// parse and validate file and post parameters
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	// check file type, detectcontenttype only needs the first 512 bytes
	detectedFileType := http.DetectContentType(fileBytes)
	switch detectedFileType {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
		break
	default:
		return errors.New("unsupport")
	}

	// Check file size
	maxFileSizeCnf := viper.Get("server.max_size")
	maxFileSize := int64(maxFileSizeCnf.(int) * 1024 * 1024)
	if maxFileSize < file.Size {
		return errors.New("File is too big")
	}

	validator.Data = fileBytes
	validator.Name = file.Filename
	validator.Path = file.Filename
	validator.Size = int(file.Size)

	return nil
}
