package responses

import (
	"math"

	"github.com/google/uuid"
)

//import "github.com/gofiber/fiber/v2"

type BaseResponse struct {
	Status  int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"result"`
}

type PaginationResponse struct {
	RecordCount int `json:"recordCount" bson:"recordCount"`
	PageCount   int `json:"pageCount" bson:"pageCount"`
	CurrentPage int `json:"currentPage" bson:"currentPage"`
	PageSize    int `json:"pageSize" bson:"pageSize"`
	Records     any `json:"records" bson:"records"`
}

type UuidComboboxElement struct {
	Uuid uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

type IntComboboxElement struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func NewBaseResponse(status int, message string, data any) BaseResponse {
	return BaseResponse{Status: status, Message: message, Data: data}
}

func NewPaginationResponse(recordCount int, currentPage int, pageSize int, records any) PaginationResponse {
	return PaginationResponse{
		RecordCount: recordCount,
		PageCount:   calculatePageCount(recordCount, pageSize),
		CurrentPage: currentPage,
		PageSize:    pageSize,
		Records:     records,
	}
}

func calculatePageCount(recordCount int, pageSize int) int {
	pageCount := float64(recordCount) / float64(pageSize)
	_pageCount := int(math.Ceil(pageCount))
	return _pageCount
}
