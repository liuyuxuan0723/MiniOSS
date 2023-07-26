package router

import (
	"awesomeProject/handler"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/index", handler.Index)
	r.GET("/file/:id", handler.GetFileInfo)
	r.GET("/file/download/:id", handler.DownloadFile)
	r.POST("/file/upload", handler.UploadFile)
	r.GET("/file/update", handler.UpdateFile)
	r.GET("/file/delete/:id", handler.DeleteFile)

	return r
}
