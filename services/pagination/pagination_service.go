package pagination

import (
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type PaginatedResponse struct {
	Items              []render.Renderer `json:"items"`
	TotalCount         int               `json:"total_count"`
	PageCount          int               `json:"page_count"`
	NextPageOffset     int               `json:"next_page_offset"`
	PreviousPageOffset int               `json:"previous_page_offset"`
}

const PAGINATION_MAX_LIMIT = 1000
const PAGINATION_DEFAULT_LIMIT = 100

func (logEntry *PaginatedResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	//logEntry.Elapsed = 10
	return nil
}

type PaginationServiceInterface interface {
	GetRequestOffsetAndLimit() (int, int)
	ApplyQueryOffsetAndLimit(query QueryOffsetAndLimit)
	PreparePaginatedResponse(items []render.Renderer, totalCount int) PaginatedResponse
}

type QueryOffsetAndLimit interface {
	Offset(offset int) *gorm.DB
	Limit(limit int) *gorm.DB
}

type PaginationService struct {
	Request *http.Request
}

func (service PaginationService) GetRequestOffsetAndLimit() (int, int) {
	offset, err := strconv.Atoi(service.Request.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(service.Request.URL.Query().Get("limit"))
	if err != nil {
		limit = PAGINATION_DEFAULT_LIMIT
	}

	return offset, limit
}

func (service PaginationService) ApplyQueryOffsetAndLimit(query QueryOffsetAndLimit) {
	offset, limit := service.GetRequestOffsetAndLimit()

	if offset < 0 {
		offset = 0
	}

	if limit < 1 {
		limit = PAGINATION_DEFAULT_LIMIT
	} else if limit > 100 {
		limit = PAGINATION_MAX_LIMIT
	}

	query.Offset(offset)
	query.Limit(limit)
}

func (service PaginationService) PreparePaginatedResponse(items []render.Renderer, totalCount int) PaginatedResponse {
	offset, limit := service.GetRequestOffsetAndLimit()

	nextPageOffset := computeNextPageOffset(offset, limit, totalCount)
	previousPageOffset := computePreviousPageOffset(offset, limit)

	paginationResponse := &PaginatedResponse{
		Items:              items,
		TotalCount:         totalCount,
		PageCount:          len(items),
		NextPageOffset:     nextPageOffset,
		PreviousPageOffset: previousPageOffset,
	}

	return *paginationResponse
}

func computeNextPageOffset(offset int, limit int, totalCount int) int {
	if offset+limit <= totalCount {
		return offset + limit
	}

	return offset
}

func computePreviousPageOffset(offset int, limit int) int {
	if offset > limit {
		return offset - limit
	}

	return 0
}
