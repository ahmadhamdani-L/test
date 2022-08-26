package upload

import (
	"fmt"
	"io"
	"kafka-go-getting-started/internal/abstraction"
	"kafka-go-getting-started/internal/dto"
	"kafka-go-getting-started/internal/factory"
	// "kafka-go-getting-started/internal/model"
	res "kafka-go-getting-started/pkg/util/response"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
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
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]
	cc := c.(*abstraction.Context)

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		xlsx, err := excelize.OpenFile(file.Filename)
		if err != nil {
			fmt.Println(err)
		}
		rows := xlsx.GetRows("Sheet One")
		i := 2
		for _, row := range rows {

			if i > 1 {
				payload := new(dto.UploadCreateRequest)
				var str []string
				for _, colCell := range row {
					str = append(str, colCell)
					fmt.Println("str append is :", str)
				}
				payload.NamaUpload = str[1]
				payload.NoUpload = str[2]
				if err := c.Bind(payload); err != nil {
					return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
				}
				if err := c.Validate(payload); err != nil {
					return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
				}

				result, err := h.service.Create(cc, payload)
				if err != nil {
					return res.ErrorResponse(err).Send(c)
				}
				return res.SuccessResponse(result).Send(c)
			}
			i++

		}

	}
	return nil

}
