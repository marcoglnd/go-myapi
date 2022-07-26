package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/go-myapi/internal/webapp/domain"
	mocks "github.com/marcoglnd/go-myapi/internal/webapp/mocks/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetData(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockReturn := domain.DataResponse{
			Data: "hello",
		}

		mockService := mocks.NewWebappService(t)

		payload, err := json.Marshal(mockReturn)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/myapi", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		webappController := WebappController{webapp: mockService}

		engine.GET("/myapi", webappController.GetData())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		mockService.AssertExpectations(t)
	})
}

func TestGetCryptoById(t *testing.T) {
	mockCurrentPrice := domain.CurrentPrice{
		Usd: 21914,
	}

	mockMarketData := domain.MarketData{
		CurrentPrice: mockCurrentPrice,
	}

	mockCrypto := domain.CryptoResponse{
		ID:         "bitcoin",
		Symbol:     "btc",
		MarketData: mockMarketData,
		Partial:    false,
	}

	type expected struct {
		status int
		err    assert.ValueAssertionFunc
	}

	tests := []struct {
		name       string
		serviceRes error
		want       expected
		method     string
		PATH       string
		route      string
	}{
		{
			name:       "success",
			serviceRes: nil,
			want:       expected{status: http.StatusOK, err: assert.Nil},
			method:     http.MethodGet,
			PATH:       fmt.Sprintf("/api/v1/crypto/%v", mockCrypto.ID),
			route:      "/api/v1/crypto/:id",
		},
		{
			name:       "partial content",
			serviceRes: errors.New("bad request"),
			want:       expected{status: http.StatusPartialContent, err: assert.Nil},
			method:     http.MethodGet,
			PATH:       fmt.Sprintf("/api/v1/crypto/%v", "aaa"),
			route:      "/api/v1/crypto/:id",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockService := mocks.NewWebappService(t)
			mockService.On("GetCryptoById",
				mock.AnythingOfType("string"),
			).Return(mockCrypto, tt.serviceRes).Once()

			payload, err := json.Marshal(mockCrypto)
			assert.NoError(t, err)

			req := httptest.NewRequest(tt.method, tt.PATH, bytes.NewBuffer(payload))
			rec := httptest.NewRecorder()

			_, engine := gin.CreateTestContext(rec)

			webappController := WebappController{webapp: mockService}

			engine.GET(tt.route, webappController.GetCryptoById())

			engine.ServeHTTP(rec, req)

			assert.Equal(t, tt.want.status, rec.Code)

			mockService.AssertExpectations(t)
		})
	}
}
