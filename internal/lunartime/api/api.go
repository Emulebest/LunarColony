package api

import (
	"LunarColony/internal/lunartime/service"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type LunarTimeController struct {
	service service.LunarTimeService
}

func (l *LunarTimeController) GetLunarTime(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	utc := &UTCParam{}
	err := json.Unmarshal(body, utc)
	if err != nil {
		zap.S().Errorf("Couldn't process request: %s", string(body))
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}
	lst := l.service.GetLunarTime(utc.UTC)
	if err = json.NewEncoder(ctx).Encode(lst); err != nil {
		zap.S().Errorf("Couldn't serialize LST time: %+v", lst)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
}

func NewLunarTimeController(service service.LunarTimeService) *LunarTimeController {
	return &LunarTimeController{
		service: service,
	}
}