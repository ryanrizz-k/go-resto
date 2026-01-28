package main

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rynr00/go-resto/internal/database"
	"github.com/rynr00/go-resto/internal/delivery/rest"
	"github.com/rynr00/go-resto/internal/logger"
	mRepo "github.com/rynr00/go-resto/internal/repository/menu"
	oRepo "github.com/rynr00/go-resto/internal/repository/order"
	uRepo "github.com/rynr00/go-resto/internal/repository/user"
	"github.com/rynr00/go-resto/internal/tracing"
	rUsecase "github.com/rynr00/go-resto/internal/usecase/resto"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=Test1234 dbname=go_resto sslmode=disable"
)

func main() {
	logger.Init()
	tracing.Init("http://localhost:14268/api/traces")
	e := echo.New()

	db := database.GetDB(dbAddress)
	secret := "AES256KEY-32Characters1234567890"
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32, signKey, 60*time.Second)
	if err != nil {
		panic(err)
	}
	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo, userRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadMiddleware(e)
	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":14045"))
}
