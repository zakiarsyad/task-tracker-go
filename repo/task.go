package repo

import (
	"tracker/model"

	"gorm.io/gorm"
)

type Tasks struct {
	conn *gorm.DB
}

func NewTasks(conn *gorm.DB) Tasks {
	return Tasks{conn}
}

func (t Tasks) Create(task model.Task) error {
	err := t.conn.Create(&task).Error
	if err != nil {
		return err
	}

	return nil
}
