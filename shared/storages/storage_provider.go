package storages

import (
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/storages/task"
	"github.com/jinzhu/gorm"
)

type IStorageProvider interface {
	Task() task.ITaskStorage
}

type StorageProvider struct {
	task task.ITaskStorage

	DB *gorm.DB
}

func NewStorageProvider(db *gorm.DB) IStorageProvider {
	return &StorageProvider{
		DB: db,
	}
}

func (s *StorageProvider) Task() task.ITaskStorage {
	if s.task == nil {
		s.task = task.NewTaskStorage(s.DB)
	}

	return s.task
}