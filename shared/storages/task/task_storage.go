package task

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type ITaskStorage interface {
	Add(task *Task) error
	FindById(id int) (*Task, error)
	UpdateById(id int, taskUpdates *Task) error
	DeleteById(id int) error
}

type TaskStorage struct {
	DB *gorm.DB
}

func NewTaskStorage(db *gorm.DB) ITaskStorage {
	return &TaskStorage{
		DB: db,
	}
}

func (ts *TaskStorage) Add(task *Task) error {
	db := ts.DB.Create(task)
	if db.Error != nil {
		return fmt.Errorf("En arror occured while adding new task %v", db.Error)
	}

	return nil
}

func (ts *TaskStorage) FindById(id int) (*Task, error) {
	task := &Task{}
	db := ts.DB.First(task, id)

	if db.RecordNotFound() {
		return task, ErrTaskNotFound
	}

	if db.Error != nil {
		return task, db.Error
	}

	return task, nil
}

func (ts *TaskStorage) UpdateById(id int, taskUpdates *Task) error {
	task := &Task{}

	db := ts.DB.First(&task, id).Update(taskUpdates)

	if db.RecordNotFound() {
		return ErrTaskNotFound
	}

	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (ts *TaskStorage) DeleteById(id int) error {
	task := &Task{}
	db := ts.DB.First(task, id).Update("is_deleted", true)

	if db.RecordNotFound() {
		return ErrTaskNotFound
	}

	if db.Error != nil {
		return db.Error
	}

	return nil
}
