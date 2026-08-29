package gomiddleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/sirupsen/logrus"
)

func EchoSetRequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		requestID := c.Request().Header.Get("X-Request-ID")

		if requestID == "" {
			uuidRequestID, err := uuid.NewRandom()

			if err != nil {
				logrus.WithFields(logrus.Fields{
					"tag":   "GoMiddleware.RequestID.EchoSetRequestID.01",
					"error": err.Error(),
				}).Error("failed to generate uuid for request id")

				return err
			}

			requestID = uuidRequestID.String()
		}

		c.Set("RequestID", &requestID)

		c.Response().Header().Set("X-Request-ID", requestID)

		return next(c)
	}
}
