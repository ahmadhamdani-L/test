package dto

import (
	"kafka-go-getting-started/internal/abstraction"
	"kafka-go-getting-started/internal/model"
	res "kafka-go-getting-started/pkg/util/response"
)

// Get
type UploadGetRequest struct {
	abstraction.Pagination
	model.UploadFilterModel
}
type UploadGetResponse struct {
	Datas          []model.UploadEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type UploadGetResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data []model.UploadEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type UploadGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type UploadGetByIDResponse struct {
	model.UploadEntityModel
}
type UploadGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta           `json:"meta"`
		Data UploadGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type UploadCreateRequest struct {
	model.UploadEntity
	// juz_id string `json:"juz_id"`
}
type UploadCreateResponse struct {
	model.UploadEntityModel
}
type UploadCreateResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data UploadCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type UploadUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.UploadEntity
}
type UploadUpdateResponse struct {
	model.UploadEntityModel
}
type UploadUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data UploadUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type UploadDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type UploadDeleteResponse struct {
	model.UploadEntityModel
}
type UploadDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta          `json:"meta"`
		Data UploadDeleteResponse `json:"data"`
	} `json:"body"`
}
