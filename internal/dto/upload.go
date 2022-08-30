package dto

import (
	// "kafka-go-getting-started/internal/abstraction"
	"kafka-go-getting-started/internal/model"
	res "kafka-go-getting-started/pkg/util/response"
)

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

