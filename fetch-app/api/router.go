package api

import (
	"monorepo/api/middleware"
	"monorepo/v1/commodity"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	commodityController *commodity.Controller,
) {
	// commodity
	commodityV1 := e.Group("v1/commodity")
	commodityV1.Use(middleware.JWTMiddleware())
	commodityV1.GET("", commodityController.GetAllCommodity)

	commodityReportV1 := e.Group("v1/report")
	commodityReportV1.Use(middleware.JWTMiddlewareAdmin())
	commodityReportV1.GET("/summary", commodityController.GetAllCommodityByProvinsi)

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
