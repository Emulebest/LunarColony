package service_test

import (
	"LunarColony/internal/lunartime/model"
	"LunarColony/internal/lunartime/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockLunarTimeService struct {
}

func NewMockLunarTimeService() service.LunarTimeService {
	return &mockLunarTimeService{}
}

func (l *mockLunarTimeService) GetLunarTime(UTC string) model.LunarTime {
	return model.LunarTime{
		LST: UTC,
	}
}

func TestGetLunarTime(t *testing.T) {
	mock := NewMockLunarTimeService()
	lst := mock.GetLunarTime("123")
	assert.Equal(t, "123", lst.LST)
}