package internal

import (
	file "govision/internal/modules/file"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.POST("/image/upload", file.UploadFileImage)
}
