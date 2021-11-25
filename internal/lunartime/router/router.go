package router

import (
	"LunarColony/internal/lunartime/api"
	"github.com/fasthttp/router"
)

type LunarTimeRouter struct {
	controller *api.LunarTimeController
	router *router.Group
}

func NewLunarTimeRouter(r *router.Group, controller *api.LunarTimeController) *LunarTimeRouter {
	lunarRouter := &LunarTimeRouter{
		controller: controller,
		router:     r,
	}
	return lunarRouter
}

func (l *LunarTimeRouter) Register() {
	l.router.POST("/convert_time", l.controller.GetLunarTime)
}
