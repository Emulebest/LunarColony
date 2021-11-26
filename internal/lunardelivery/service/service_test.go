package service_test

import (
	"LunarColony/internal/lunardelivery/model"
	"LunarColony/internal/lunardelivery/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockLunarTimeService struct {
}

func NewMockLunarDeliveryService() service.LunarDeliveryService {
	return &mockLunarTimeService{}
}

func (l *mockLunarTimeService) GetDeliveryTime(UTC string) (model.LunarTime, error) {
	return model.LunarTime{
		LST: UTC,
	}, nil
}

func TestGetDeliveryTimeMock(t *testing.T) {
	mock := NewMockLunarDeliveryService()
	lst, err := mock.GetDeliveryTime("123")
	assert.NoError(t, err)
	assert.Equal(t, "123", lst.LST)
}

func TestGetDeliveryTime(t *testing.T) {
	srvc := service.NewLunarDeliveryService()
	lst, err := srvc.GetDeliveryTime("2021-11-26 11:28:45 UTC")
	assert.NoError(t, err)
	assert.Equal(t, "54-12-20 delta 08:49:43", lst.LST)
}
