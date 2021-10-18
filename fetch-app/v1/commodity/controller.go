package commodity

import (
	"net/http"

	businessCommodity "monorepo/business/commodity"
	common "monorepo/common"
	templateResponse "monorepo/v1/commodity/response"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	commodityService businessCommodity.Service
}

func NewController(commodityService businessCommodity.Service) *Controller {
	return &Controller{
		commodityService,
	}
}

func (controller Controller) GetAllCommodity(c echo.Context) error {
	commodities, err := controller.commodityService.FindAllCommodity()
	if err != nil {
		return c.JSON(http.StatusNotFound, "Commodities Not Found")
	}

	response := common.NewSuccessResponse(templateResponse.NewCommodityDataResponse(commodities))

	return c.JSON(http.StatusOK, response)
}

func (controller Controller) GetAllCommodityByProvinsi(c echo.Context) error {
	commodities, err := controller.commodityService.FindAllCommodityByProvinsi()
	if err != nil {
		return c.JSON(http.StatusNotFound, "Commodities Not Found")
	}

	response := common.NewSuccessResponse(templateResponse.NewCommodityDataReportResponse(commodities))

	return c.JSON(http.StatusOK, response)
}
