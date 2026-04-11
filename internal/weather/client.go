package weather

import (
	"github.com/avaswani-build/fair-winds-api/internal/domain"
)

type Client interface {
	GetForecast(location string) (domain.Forecast, error)
}

type MockClient struct{}

func (m MockClient) GetForecast(location string) (domain.Forecast, error) {
	return domain.Forecast{
		Location:   location,
		WindAvgKts: 12,
		GustKts:    18,
		WindDir:    "SW",
	}, nil
}
