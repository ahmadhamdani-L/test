package repository

import (
	"fmt"
	"io"
	"kafka-go-getting-started/internal/abstraction"
	"kafka-go-getting-started/internal/model"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

type Upload interface {
	Create(ctx *abstraction.Context, e *model.UploadEntity) (*model.UploadEntityModel, error)
	Read(ctx *abstraction.Context, e *model.UploadEntity) (*model.UploadEntityModel, error)
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

	// err = conn.Model(data).First(&data).
	// 	WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *upload) Read(ctx *abstraction.Context, e *model.UploadEntity) (*model.UploadEntityModel, error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}
	conn := r.CheckTrx(ctx)
	var data model.UploadEntityModel
	files := form.File["files"]
	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(file.Filename)
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return nil, err
		}

		xlsx, err := excelize.OpenFile(file.Filename)
		if err != nil {
			fmt.Println(err)
		}

		rows:= xlsx.GetRows("Sheet One")
		i := 1
		for _, row := range rows {

			if i > 1 {

				var str []string
				for _, colCell := range row {
					str = append(str, colCell)
				}

				data.NamaUpload = str[0]
				data.NoUpload = str[1]

				sqlStatement := `INSERT INTO uploads (nama_upload, no_upload) 
				VALUES ($1, $2)`
				_= conn.Exec(sqlStatement, data.NamaUpload, data.NoUpload )
				if err != nil {
					panic(err)
				} else {
					fmt.Println("\nRow inserted successfully!")
				}

			}
			i++
		}
		
	}
	return &data, nil
}