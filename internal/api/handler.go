package api

import (
	"net/http"
	"strconv"

	"github.com/avaswani-build/fair-winds-api/internal/domain"
	"github.com/avaswani-build/fair-winds-api/internal/service"
	"github.com/gin-gonic/gin"
)

type WeatherClient interface {
	GetForecast(lat, lng float64) (domain.Forecast, error)
}

type Handler struct {
	WeatherClient WeatherClient
}

type SummaryResponse struct {
	Forecast       domain.Forecast       `json:"forecast"`
	Recommendation domain.Recommendation `json:"recommendation"`
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func SummaryMock(c *gin.Context) {
	forecast := domain.Forecast{
		Location:   "NY Harbor",
		WindAvgKts: 12,
		GustKts:    18,
		WindDir:    "SW",
	}

	recommendation := service.Recommend(forecast)

	c.JSON(http.StatusOK, SummaryResponse{
		Forecast:       forecast,
		Recommendation: recommendation,
	})
}

func (h *Handler) Summary(c *gin.Context) {
	latStr := c.Query("lat")
	lngStr := c.Query("lng")

	if latStr == "" || lngStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing required query params: lat and lng",
		})
		return
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid lat value",
		})
		return
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid lng value",
		})
		return
	}

	forecast, err := h.WeatherClient.GetForecast(lat, lng)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch forecast",
		})
		return
	}

	recommendation := service.Recommend(forecast)

	c.JSON(http.StatusOK, SummaryResponse{
		Forecast:       forecast,
		Recommendation: recommendation,
	})
}
