package model

import (
	"github.com/jinzhu/gorm"
	"moges/storage"
)

type Photo struct {
	gorm.Model
	Name string
	Path string
	Size int
}


func (model *Photo) Save() error {
	db := storage.GetDB()
	err := db.Create(&model).Error
	return err
}