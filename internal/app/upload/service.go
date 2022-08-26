package upload

import (
	"kafka-go-getting-started/internal/abstraction"
	"kafka-go-getting-started/internal/dto"
	"kafka-go-getting-started/internal/factory"
	"kafka-go-getting-started/internal/model"
	"kafka-go-getting-started/internal/repository"
	res "kafka-go-getting-started/pkg/util/response"
	"kafka-go-getting-started/pkg/util/trxmanager"

	"gorm.io/gorm"
)

type Service interface {
	Create(ctx *abstraction.Context, payload *dto.UploadCreateRequest) (*dto.UploadCreateResponse, error)
}

type service struct {
	Repository repository.Upload
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.UploadRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.UploadCreateRequest) (*dto.UploadCreateResponse, error) {
	var result *dto.UploadCreateResponse
	var data *model.UploadEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		// data.Context = ctx

		// data.UploadEntity = payload.UploadEntity
		data, err = s.Repository.Create(ctx, &payload.UploadEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}
	result = &dto.UploadCreateResponse{
		UploadEntityModel: *data,
	}

	return result, nil
}