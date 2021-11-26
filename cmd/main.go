package main

import (
	"LunarColony/internal/lunardelivery/api"
	"LunarColony/internal/lunardelivery/router"
	"LunarColony/internal/lunardelivery/service"
	"fmt"
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
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Println("Couldn't flush logging buffer")
		}
	}(logger) // flushes buffer, if any
	sugar := logger.Sugar()
	r := fasthttprouter.New()
	apiGroup := r.Group("/api/v1")
	{
		sugar.Info("Registering Lunar services")
		lunarGroup := apiGroup.Group("/lunar")
		lunarService := service.NewLunarDeliveryService()
		lunarController := api.NewLunarDeliveryController(lunarService)
		lunarRouter := router.NewLunarDeliveryRouter(lunarGroup, lunarController)
		lunarRouter.Register()
		sugar.Info("Finished registering Lunar services")
	}
	sugar.Info("Starting server...")
	sugar.Fatal(fasthttp.ListenAndServe(HOST+PORT, r.Handler))
}
