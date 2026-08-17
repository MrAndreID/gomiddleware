package gomiddleware

import (
	"github.com/labstack/echo/v5"
)

func EchoSetRouteList(routes echo.Routes) map[string]map[string]string {
	var routeList map[string]map[string]string = make(map[string]map[string]string)

	for _, element := range routes {
		if element.Method == echo.RouteNotFound {
			continue
		}

		if _, ok := routeList[element.Path]; !ok {
			routeList[element.Path] = make(map[string]string)
		}

		routeList[element.Path][element.Method] = element.Name
	}

	return routeList
}
