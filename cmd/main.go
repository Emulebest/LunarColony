package main

import (
	_ "LunarColony/docs"
	"LunarColony/internal/lunardelivery/api"
	"LunarColony/internal/lunardelivery/router"
	"LunarColony/internal/lunardelivery/service"
	"fmt"
	fasthttprouter "github.com/fasthttp/router"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.uber.org/zap"
)

const (
	HOST = "0.0.0.0"
	PORT = ":8080"
)

// @title           Swagger Lunar API
// @version         1.0
// @description     This is a microservice serving lunar mission
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			// No reason to panic here but logging looks broken in this case
			fmt.Println("Couldn't flush logging buffer")
		}
	}(logger) // flushes buffer, if any
	sugar := logger.Sugar()

	// Define router with Swagger
	r := fasthttprouter.New()
	r.GET("/docs/{filepath:*}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))

	// Define api group
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
