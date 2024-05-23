package models

import "gorm.io/gorm"

type ID = int

type baseModel struct {
	model *gorm.DB
}

func (m *baseModel) StartSession() {
	m.model = m.model.Session(&gorm.Session{})
}

var models = []interface{}{
	&User{},
	&Invoice{},
}
