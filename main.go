package main

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	_getCurrentTimeHandler "nextal_test/domains/GetCurrentTime/delivery/http"
	_getCurrentTimeUseCase "nextal_test/domains/GetCurrentTime/usecase"
	"nextal_test/logger"
	"nextal_test/models"
	"time"
)

var ech *echo.Echo

type CustomValidator struct {
	validator *validator.Validate
}

func main() {

	echoGroup := models.EchoGroup{
		API: ech.Group(""),
	}

	customValidator := validator.New()

	ech.Validator = &CustomValidator{validator: customValidator}

	_getCurrentTimeUseCase := _getCurrentTimeUseCase.NewGetCurrentTimeUseCase()
	_getCurrentTimeHandler.NewGetCurrentTimeHandler(echoGroup, _getCurrentTimeUseCase)

	ech.GET("/ping", ping)

	err := ech.Start(":" + "8080")

	if err != nil {
		logger.Make(nil, nil).Debug(err)
	}
}

func init() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
	ech = echo.New()
	ech.Debug = true

	logger.Init()
}

func ping(echTx echo.Context) error {

	response := map[string]interface{}{
		"status": "online",
		"month":  "Server Actived!!",
	}

	return echTx.JSON(http.StatusOK, response)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
