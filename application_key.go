package gomiddleware

import (
	"errors"

	"github.com/labstack/echo/v5"
	"github.com/sirupsen/logrus"
)

func EchoCheckApplicationKey(key string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			var tag string = "GoMiddleware.ApplicationKey.EchoCheckApplicationKey."

			applicationKey := c.Request().Header.Get("X-App-Key")

			if applicationKey == "" {
				logrus.WithFields(logrus.Fields{
					"tag": tag + "01",
				}).Error("application key not found")

				return errors.New("APPLICATION_KEY_NOT_FOUND")
			}

			if applicationKey != key {
				logrus.WithFields(logrus.Fields{
					"tag": tag + "02",
				}).Error("application key not match")

				return errors.New("APPLICATION_KEY_NOT_MATCH")
			}

			return next(c)
		}
	}
}
