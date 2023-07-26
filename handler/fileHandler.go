package handler

import (
	"awesomeProject/model"
	"awesomeProject/repo"
	"awesomeProject/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func GetFileInfo(c *gin.Context) {
	id := c.Param("id")
	fileMeta, _ := repo.QueryFileById(id)
	c.JSON(http.StatusOK, fileMeta)
}

func IsFileExists(c *gin.Context) {
	// 检查temp文件夹中是否存在md5，或者接入redis后可查
	c.JSON(http.StatusOK, new(Response))
}

func DownloadFile(c *gin.Context) {
	// 假设已经有一个 os.File 对象 file
	fileMeta, err := repo.QueryFileById(c.Param("id"))
	if err != nil {
		c.String(http.StatusInternalServerError, "file not found")
		return
	}
	file, err := os.Open(fileMeta.Location)
	defer file.Close()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error opening file")
		return
	}

	// 设置响应头，指定文件名和 Content-Disposition
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileMeta.FileName))

	// 将文件内容写入响应
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error writing file to response")
		return
	}
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, new(Response).ERROR().SetMsg("Error writing file to response"))
		return
	}
	md5, err := util.MD5(file)

	fileMeta := &model.FileMeta{
		MD5:      md5,
		FileName: file.Filename,
		FileSize: file.Size,
		Location: model.FileStorage + "/" + file.Filename,
	}
	err = repo.CreateFile(fileMeta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, new(Response).ERROR().SetMsg("filed to create file"))
	}
	err = c.SaveUploadedFile(file, fileMeta.Location)
	c.JSON(http.StatusOK, new(Response).OK())
}

func DeleteFile(c *gin.Context) {
	err := repo.DeleteFile(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, new(Response).ERROR())
		return
	}
	c.JSON(http.StatusOK, new(Response))
}

func UpdateFile(c *gin.Context) {
	file := &model.FileMeta{
		FileName: c.Param("filename"),
		// 虚拟的文件目录 并不真实存在
		Location: c.Param("location"),
	}
	err := repo.UpdateFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, new(Response).ERROR())
		return
	}
	c.JSON(http.StatusOK, new(Response))
}
