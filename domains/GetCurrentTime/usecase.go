package GetCurrentTime

import (
	"github.com/labstack/echo"
	"nextal_test/models"
)

type Usecase interface {
	UGetCurrentTime(c echo.Context, params models.ParameterTimeZone) (interface{}, error)
}
