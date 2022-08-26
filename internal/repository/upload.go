package repository

import (
	"kafka-go-getting-started/internal/abstraction"
	"kafka-go-getting-started/internal/model"

	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

type Upload interface {
	Create(ctx *abstraction.Context, e *model.UploadEntity) (*model.UploadEntityModel, error)
}

type upload struct {
	abstraction.Repository
}

func NewUpload(db *gorm.DB) *upload {
	return &upload{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *upload) Create(ctx *abstraction.Context, e *model.UploadEntity) (*model.UploadEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.UploadEntityModel
	data.UploadEntity = *e
	var Nama = ctx.Auth.Name
	data.CreatedBy = Nama
	err := conn.Create(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	err = conn.Model(data).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
