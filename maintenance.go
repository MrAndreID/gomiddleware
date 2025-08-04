package gomiddleware

import (
	"errors"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func EchoSetMaintenanceMode(fileName string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if _, err := os.Stat("maintenance.flag"); err == nil {
				logrus.WithFields(logrus.Fields{
					"tag": "GoMiddleware.Maintenance.EchoSetMaintenanceMode.01",
				}).Error("maintenance mode")

				return errors.New("MAINTENANCE")
			}

			return next(c)
		}
	}
}
