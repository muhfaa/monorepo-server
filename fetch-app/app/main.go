package main

import (
	"context"
	"database/sql"
	"fmt"
	stdLog "log"
	"os"
	"os/signal"
	"time"

	efisheryApp "monorepo/api"
	commodityAPIV1Module "monorepo/v1/commodity"

	commodity "monorepo/business/commodity"
	commodityModule "monorepo/modules/commodity"

	"monorepo/config"
	sqlite "monorepo/modules/sqllite"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var banner = `
███████╗███████╗██╗░██████╗██╗░░██╗███████╗██████╗░██╗░░░██╗
██╔════╝██╔════╝██║██╔════╝██║░░██║██╔════╝██╔══██╗╚██╗░██╔╝
█████╗░░█████╗░░██║╚█████╗░███████║█████╗░░██████╔╝░╚████╔╝░
██╔══╝░░██╔══╝░░██║░╚═══██╗██╔══██║██╔══╝░░██╔══██╗░░╚██╔╝░░
███████╗██║░░░░░██║██████╔╝██║░░██║███████╗██║░░██║░░░██║░░░
╚══════╝╚═╝░░░░░╚═╝╚═════╝░╚═╝░░╚═╝╚══════╝╚═╝░░╚═╝░░░╚═╝░░░`

func newDatabaseConnection() *sql.DB {

	db, err := sqlite.NewDatabaseConnection(config.GetConfig().DbName)
	if err != nil {
		panic(err)
	}

	return db

}

func newCommodityService() commodity.Service {
	// commodity repo
	commodityRepository := commodityModule.NewHTTPRequestRepository()
	// commodity service
	commodityService := commodity.NewService(commodityRepository)

	return commodityService
}

func main() {
	stdLog.Println(banner)

	// initialize database connection based on given config
	dbCon := newDatabaseConnection()

	commodityService := newCommodityService()

	commodityAPIController := commodityAPIV1Module.NewController(commodityService)

	e := echo.New()
	e.HideBanner = true

	efisheryApp.RegisterPath(
		e,
		commodityAPIController,
	)

	// run server
	go func() {
		// address := fmt.Sprintf("0.0.0.0:%d", config.GetConfig().Port)
		port := 8000
		address := fmt.Sprintf("0.0.0.0:%d", port)

		if err := e.Start(address); err != nil {
			log.Error("failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// db close
	defer dbCon.Close()

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Error("failed to grafefully shutting down server", err)
	}

}
