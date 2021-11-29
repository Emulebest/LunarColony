package api

import (
	"LunarColony/internal/lunardelivery/service"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

// LunarDeliveryController is responsible for processing request and applying the business logic
type LunarDeliveryController struct {
	service service.LunarDeliveryService
}

// GetDeliveryTime works with request parameters and passes on the the service layer
// @Summary Calculates the delivery time
// @Description Calculates the delivery time from Earth to Moon based on the UTC and internal parameters
// @Accept json
// @Produce json
// @Param utc body UTCParam true "UTC input"
// @Success 200 {object} model.LunarTime "Lunar Standard Time"
// @Router /lunar/delivery_time [post]
func (l *LunarDeliveryController) GetDeliveryTime(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	utc := &UTCParam{}
	err := json.Unmarshal(body, utc)
	if err != nil {
		zap.S().Errorf("Couldn't process request: %s", string(body))
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}
	zap.S().Debug("Getting delivery time")
	lst, err := l.service.GetDeliveryTime(utc.UTC)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}
	zap.S().Debug("Finished getting delivery time")
	if err = json.NewEncoder(ctx).Encode(lst); err != nil {
		zap.S().Errorf("Couldn't serialize LST time: %+v", lst)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
}

// NewLunarDeliveryController constructs a controller
func NewLunarDeliveryController(service service.LunarDeliveryService) *LunarDeliveryController {
	return &LunarDeliveryController{
		service: service,
	}
}
