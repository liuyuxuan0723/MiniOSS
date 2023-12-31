package repo

import (
	"MiniOSS/internal/model"
	"MiniOSS/pkgs"
	"fmt"
)

func QueryFileById(id string) (*model.FileMeta, error) {
	var file = &model.FileMeta{}
	err := pkgs.DB.Where("id=?", id).Take(file).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query file: %w", err)
	}
	return file, nil
}

func QueryFileByCond(name string, location string) ([]*model.FileMeta, error) {
	files := make([]*model.FileMeta, 10)
	err := pkgs.DB.Where("file_name = ? AND location LIKE ?", name, location+"%").Take(files).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query file: %w", err)
	}
	return files, nil
}

func QueryByMD5(md5 string) (*model.FileMeta, error) {
	var file = &model.FileMeta{}
	err := pkgs.DB.Where("md5=?", md5).Take(file).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query file: %w", err)
	}
	return file, nil
}

func CreateFile(file *model.FileMeta) error {
	if err := pkgs.DB.Create(file).Error; err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	return nil
}

func UpdateFile(file *model.FileMeta) error {
	if err := pkgs.DB.Updates(file).Error; err != nil {
		return fmt.Errorf("failed to update file: %w", err)
	}
	return nil
}

func DeleteFile(id string) error {
	if err := pkgs.DB.Delete(&model.FileMeta{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}
