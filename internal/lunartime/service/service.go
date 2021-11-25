package service

import "LunarColony/internal/lunartime/model"

type LunarTimeService interface {
	GetLunarTime(UTC string) model.LunarTime
}

type lunarTimeService struct {
}

// TODO: change to real impl
func (l *lunarTimeService) GetLunarTime(UTC string) model.LunarTime {
	return model.LunarTime{
		LST: UTC,
	}
}

func NewLunarTimeService() LunarTimeService {
	return &lunarTimeService{}
}
