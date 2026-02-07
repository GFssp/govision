package file

import (
	"mime/multipart"
)

type UploadRequest struct {
	File multipart.File
}

type ImgBBResponse struct {
	Data struct {
		URL string `json:"url"`
	}
}
