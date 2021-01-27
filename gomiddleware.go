package gomiddleware

import (
	"fmt"
	"log"
	"net/http"
)

func Log(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		log.Println("Accepted")
		log.Println(request.Method + " " + request.Host + request.URL.String())

		next(response, request)

		log.Println("Closing")
		fmt.Println()
	})
}
