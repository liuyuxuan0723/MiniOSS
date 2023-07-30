package api

import (
	"MiniOSS/internal/service"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/index", service.Index)
	r.GET("/file/:id", service.GetFileInfo)
	r.GET("/file/download/:id", service.DownloadFile)
	r.POST("/file/upload", service.UploadFile)
	r.GET("/file/update", service.UpdateFile)
	r.GET("/file/delete/:id", service.DeleteFile)

	return r
}
