package router

import (
	"LunarColony/internal/lunardelivery/api"
	"github.com/fasthttp/router"
)

// LunarDeliveryRouter is responsible for routing components
type LunarDeliveryRouter struct {
	controller *api.LunarDeliveryController
	router     *router.Group
}

// NewLunarDeliveryRouter constructs a router
func NewLunarDeliveryRouter(r *router.Group, controller *api.LunarDeliveryController) *LunarDeliveryRouter {
	lunarRouter := &LunarDeliveryRouter{
		controller: controller,
		router:     r,
	}
	return lunarRouter
}

// Register registers routing patterns for the lunar delivery in the global router
func (l *LunarDeliveryRouter) Register() {
	l.router.POST("/delivery_time", l.controller.GetDeliveryTime)
}
