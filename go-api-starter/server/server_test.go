package server

import (
	"errors"
	"fmt"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerRoutes(t *testing.T) {

	tests := []struct {
		path                 string
		expectedResponseCode int
	}{
		{
			"/api/v1/",
			200,
		},
		{
			"/api/v1/error",
			500,
		},
		{
			"/api/v1/not/exist",
			404,
		},
	}

	svr := &Server{
		App:      InitFiberApplication(),
		QuoteSvc: &mockQuotter{},
	}

	svr.setup()

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.path, nil)
		res, err := svr.App.Test(req)
		assert.Nil(t, err, "Error making request")
		failedPath := fmt.Sprintf("Failed for %s", test.path)
		assert.Equal(t, test.expectedResponseCode, res.StatusCode, failedPath)
	}

}

type mockQuotter struct {
	quote    string
	quoteErr error
}

func (mq *mockQuotter) GetQuote() (string, error) {
	return mq.quote, mq.quoteErr
}

func TestQuoteService(t *testing.T) {

	path := "/api/v1/quote"
	expRespCode := 200

	svr := &Server{
		App:      InitFiberApplication(),
		QuoteSvc: &mockQuotter{"Hello", nil},
	}

	svr.setup()
	req := httptest.NewRequest("GET", path, nil)
	res, err := svr.App.Test(req)
	assert.Nil(t, err, "Error making request")
	failedPath := fmt.Sprintf("Failed for %s", path)
	if res == nil {
		log.Println("ReS is nil")
	}
	assert.Equal(t, expRespCode, res.StatusCode, failedPath)
}

func TestQuoteError(t *testing.T) {
	path := "/api/v1/quote"
	expRespCode := 500

	svr := &Server{
		App:      InitFiberApplication(),
		QuoteSvc: &mockQuotter{"Hello", errors.New("Some Exception")},
	}

	svr.setup()
	req := httptest.NewRequest("GET", path, nil)
	res, err := svr.App.Test(req)
	assert.Nil(t, err, "Error making request")
	failedPath := fmt.Sprintf("Failed for %s", path)
	if res == nil {
		log.Println("ReS is nil")
	}
	assert.Equal(t, expRespCode, res.StatusCode, failedPath)
}
