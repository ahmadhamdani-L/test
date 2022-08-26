package model

import (
	"kafka-go-getting-started/internal/abstraction"
	"kafka-go-getting-started/pkg/util/date"
	"gorm.io/gorm"
)

type UploadEntity struct {
	NamaUpload string `json:"nama_upload" validate:"required" gorm:"index:idx_upload_nama_upload,unique"`
	NoUpload   string `json:"no_upload" validate:"required"`
}

type UploadFilter struct {
	NamaUpload *string `query:"nama_upload" filter:"ILIKE"`
	NoUpload   *string `query:"no_upload"`
}

type UploadEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	UploadEntity

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type UploadFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	UploadFilter
}

func (UploadEntityModel) TableName() string {
	return "uploads"
}

func (m *UploadEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}

func (m *UploadEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = *date.DateTodayLocal()
	return
}
