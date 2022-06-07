package pagination

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type MockedQuery struct {
	OffsetValue int
	LimitValue  int
}

func (query *MockedQuery) Offset(offset int) *gorm.DB {
	query.OffsetValue = offset

	return &gorm.DB{}
}

func (query *MockedQuery) Limit(limit int) *gorm.DB {
	query.LimitValue = limit

	return &gorm.DB{}
}

func TestApplyQueryOffsetAndLimit(test *testing.T) {
	expectedOffset := 10
	expectedLimit := 20

	url := fmt.Sprintf("/test?offset=%d&limit=%d", expectedOffset, expectedLimit)
	mockedRequest, err := http.NewRequest(http.MethodGet, url, bytes.NewBufferString(""))
	if err != nil {
		log.Fatalf("Error mocking request: %s", err)
	}

	paginationService := PaginationService{Request: mockedRequest}
	query := MockedQuery{}
	paginationService.ApplyQueryOffsetAndLimit(&query)

	if query.OffsetValue != expectedOffset {
		test.Fatalf("Failed asserting that the query offset value '%d' is equal to the expected value '%d'.", query.OffsetValue, expectedOffset)
	}

	if query.LimitValue != expectedLimit {
		test.Fatalf("Failed asserting that the query limit value '%d' is equal to the expected value '%d'.", query.LimitValue, expectedLimit)
	}

	test.Log("Finished testing the ApplyQueryOffsetAndLimit() method")
}

func TestGetRequestOffsetAndLimit(test *testing.T) {
	expectedOffset := 10
	expectedLimit := 20

	url := fmt.Sprintf("/test?offset=%d&limit=%d", expectedOffset, expectedLimit)
	mockedRequest, err := http.NewRequest(http.MethodGet, url, bytes.NewBufferString(""))
	if err != nil {
		log.Fatalf("Error mocking request: %s", err)
	}

	paginationService := PaginationService{Request: mockedRequest}
	computedOffset, computedLimit := paginationService.GetRequestOffsetAndLimit()

	if computedOffset != expectedOffset {
		test.Fatalf("Failed asserting that the computed offset value '%d' is equal to the expected value '%d'.", computedOffset, expectedOffset)
	}

	if computedLimit != expectedLimit {
		test.Fatalf("Failed asserting that the computed limit value '%d' is equal to the expected value '%d'.", computedLimit, expectedLimit)
	}

	test.Log("Finished testing the GetRequestOffsetAndLimit() method")
}

func TestPreparePaginatedResponse(test *testing.T) {
	expectedTotalCount := 5
	expectedOffset := 10
	expectedLimit := 20

	url := fmt.Sprintf("/test?offset=%d&limit=%d", expectedOffset, expectedLimit)
	mockedRequest, err := http.NewRequest(http.MethodGet, url, bytes.NewBufferString(""))
	if err != nil {
		log.Fatalf("Error mocking request: %s", err)
	}

	paginationService := PaginationService{Request: mockedRequest}
	var items []render.Renderer
	paginatedResponse := paginationService.PreparePaginatedResponse(items, expectedTotalCount)

	if paginatedResponse.TotalCount != expectedTotalCount {
		test.Fatalf("Failed asserting that the total count of items is '%d' (got '%d').", expectedTotalCount, paginatedResponse.TotalCount)
	}

	expectedPageCount := len(items)
	if paginatedResponse.PageCount != expectedPageCount {
		test.Fatalf("Failed asserting that the page count of items is '%d' (got '%d').", expectedPageCount, paginatedResponse.PageCount)
	}

	expectedNextPageOffset := 10
	if paginatedResponse.NextPageOffset != expectedNextPageOffset {
		test.Fatalf("Failed asserting that the page offset is '%d' (got '%d').", expectedNextPageOffset, paginatedResponse.NextPageOffset)
	}

	expectedPreviousPageOffset := 0
	if paginatedResponse.PreviousPageOffset != expectedPreviousPageOffset {
		test.Fatalf("Failed asserting that the previous page offset is '%d' (got '%d').", expectedPreviousPageOffset, paginatedResponse.PreviousPageOffset)
	}

	test.Log("Finished testing the PreparePaginatedResponse() method")
}

func TestComputeNextPageOffset(test *testing.T) {
	computedOffset := computeNextPageOffset(300, 100, 500)
	expectedOffset := 400
	if computedOffset != expectedOffset {
		test.Fatal("Failed asserting that the next page offset will be 400 with size = 100, current offset = 300, and total result = 500.")
	}

	computedOffset = computeNextPageOffset(50, 100, 150)
	expectedOffset = 150
	if computedOffset != expectedOffset {
		test.Fatalf("Failed asserting that the next page offset is '%d' (got '%d').", expectedOffset, computedOffset)
	}

	computedOffset = computeNextPageOffset(100, 100, 150)
	expectedOffset = 100
	if computedOffset != expectedOffset {
		test.Fatalf("Failed asserting that the next page offset is '%d' (got '%d').", expectedOffset, computedOffset)
	}

	test.Log("Finished testing the computeNextPageOffset() method")
}

func TestComputePreviousPageOffset(test *testing.T) {
	if computePreviousPageOffset(300, 100) != 200 {
		test.Fatal("Failed asserting that the previous page offset will be 200 with size = 100 and current offset = 200.")
	}

	if computePreviousPageOffset(50, 100) != 0 {
		test.Fatal("Failed asserting that the previous page offset will be 0 with size = 100 and current offset = 50.")
	}

	test.Log("Finished testing the computePreviousPageOffset() method")
}
