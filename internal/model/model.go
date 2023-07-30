package model

import "gorm.io/gorm"

const FileStorage = "/Users/liuyuxuan/Documents/workbench/awesomeProject/resource"
const TempStorage = "/Users/liuyuxuan/Documents/workbench/awesomeProject/temp"

type FileMeta struct {
	gorm.Model `json:"gorm_._model"`
	MD5        string `gorm:"md5,omitempty"`
	FileName   string `gorm:"file_name,omitempty"`
	FileSize   int64  `gorm:"file_size,omitempty"`
	Location   string `gorm:"location,omitempty"`
}

func (f *FileMeta) TableName() string {
	return "file_meta"
}

type User struct {
	gorm.Model `json:"gorm_._model"`
	UserName   string `gorm:"username,omitempty"`
	Password   string `gorm:"password,omitempty"`
	Role       string `gorm:"role,omitempty"`
}

type File struct {
	gorm.Model `json:"Gorm.Model"`
	FileID     string `gorm:"function_id,omitempty"`
	UserID     string `gorm:"user_id,omitempty"`
}
