package upload

import (

	"kafka-go-getting-started/internal/abstraction"
	"kafka-go-getting-started/internal/dto"
	"kafka-go-getting-started/internal/factory"

	// "kafka-go-getting-started/internal/model"
	res "kafka-go-getting-started/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service
}

var err error

func NewHandler(f *factory.Factory) *handler {
	service := NewService(f)
	return &handler{service}
}

func (h *handler) Create(c echo.Context) error {
	cc := c.(*abstraction.Context)
	payload := new(dto.UploadCreateRequest)
	result, err := h.service.Read(cc,payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}



// func (h *handler) Upload(c echo.Context) error {
// 	cc := c.(*abstraction.Context)

// 	payload := new(dto.FileHeader)
// 	if err := c.Bind(payload); err != nil {
// 		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
// 	}
// 	if err := c.Validate(payload); err != nil {
// 		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
// 	}

// 	result, err := h.service.Create(cc, payload)
// 	if err != nil {
// 		return res.ErrorResponse(err).Send(c)
// 	}

// 	return res.SuccessResponse(result).Send(c)
// }