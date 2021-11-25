package main

import (
	"LunarColony/internal/lunartime/api"
	"LunarColony/internal/lunartime/router"
	"LunarColony/internal/lunartime/service"
	fasthttprouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

const (
	HOST = ""
	PORT = ":8080"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	r := fasthttprouter.New()
	apiGroup := r.Group("/api/v1")
	{
		sugar.Info("Registering Lunar services")
		lunarGroup := apiGroup.Group("/lunartime")
		lunarService := service.NewLunarTimeService()
		lunarController := api.NewLunarTimeController(lunarService)
		lunarRouter := router.NewLunarTimeRouter(lunarGroup, lunarController)
		lunarRouter.Register()
		sugar.Info("Finished registering Lunar services")
	}
	sugar.Info("Starting server...")
	sugar.Fatal(fasthttp.ListenAndServe(HOST+PORT, r.Handler))
}
