package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lazytooo/lottery-draw/handler"
	"github.com/lazytooo/lottery-draw/repo"
	"github.com/lazytooo/lottery-draw/usecase"
	commonConfig "github.com/lazytooo/utils-center/config"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())

	configs := commonConfig.NewConfig()
	db := configs.MysqlInit()
	rds := configs.RedisInit()
	repo := repo.NewRepository(db, rds)
	ucase := usecase.NewUsecase(repo)
	handler.NewHttpHandler(e, ucase)

	e.Start(":8000")
}
