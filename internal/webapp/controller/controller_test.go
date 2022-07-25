package controller

import (
	"bytes"
	"encoding/json"
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
	t.Run("success", func(t *testing.T) {
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

		mockService := mocks.NewWebappService(t)
		mockService.On("GetCrypto",
			mock.AnythingOfType("string"),
		).Return(&mockCrypto, nil).Once()

		payload, err := json.Marshal(mockCrypto)
		assert.NoError(t, err)

		PATH := fmt.Sprintf("/crypto/%v", mockCrypto.ID)
		req := httptest.NewRequest(http.MethodGet, PATH, bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		webappController := WebappController{webapp: mockService}

		engine.GET("/crypto/:id", webappController.GetCryptoById())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		mockService.AssertExpectations(t)
	})
}
