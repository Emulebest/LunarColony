package service

import (
	"LunarColony/internal/lunardelivery/model"
	"LunarColony/internal/lunardelivery/service/converters"
	"go.uber.org/zap"
	"time"
)

const (
	// daysWarehouseToStation amount of days needed to move a package from warehouse to the space station on Earth
	daysWarehouseToStation = 1
	// daysEarthToMoon amount of days needed to transport a package from the Earth space station to the Moon
	daysEarthToMoon = 3
)

// LunarDeliveryService specifies the function required for the lunar delivery service
type LunarDeliveryService interface {
	GetDeliveryTime(UTC string) (model.LunarTime, error)
}

// lunarDeliveryService wrapper for the interface implementation
type lunarDeliveryService struct {
}

// GetDeliveryTime calculates the total delivery time needed for the package, full cycle
func (l *lunarDeliveryService) GetDeliveryTime(UTC string) (model.LunarTime, error) {
	parsedTime, err := converters.StringToUTC(UTC)
	if err != nil {
		zap.S().Errorf("Couldn't convert request to golang time.Time: %s", UTC)
		return model.LunarTime{}, err
	}
	deliveryTimeEarth := l.calculateDeliveryTime(parsedTime)
	lunarElapsedSeconds := converters.LunarElapsedSeconds(deliveryTimeEarth)
	lunarSeconds := converters.ConvertEarthToMoonSec(lunarElapsedSeconds)
	lst := converters.FormatLST(lunarSeconds)
	return model.LunarTime{
		LST: lst,
	}, nil
}

// calculateDeliveryTime calculates the delivery time in Earth days based on the constants above
func (l *lunarDeliveryService) calculateDeliveryTime(startTime time.Time) time.Time {
	return startTime.AddDate(0, 0, daysWarehouseToStation+daysEarthToMoon)
}

// NewLunarDeliveryService constructs an instance of the Delivery Service
func NewLunarDeliveryService() LunarDeliveryService {
	return &lunarDeliveryService{}
}
