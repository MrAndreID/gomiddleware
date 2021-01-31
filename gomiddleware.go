package gomiddleware

import (
	"fmt"
	"log"
	"net/http"

	helpers "github.com/MrAndreID/gohelpers"
)

func Log(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		log.Println("Accepted")
		log.Println("[" + request.Method + "] " + request.Host + request.URL.String())
		fmt.Println()

		next(response, request)

		log.Println("Closing")
		fmt.Println()
	})
}

func Acceptable(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Header["Accept"] == nil || request.Header["Accept"][0] == "" {
			helpers.HandleResponse(response, 406, "not acceptable", nil)

			return
		}

		next(response, request)
	})
}

func ClientID(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Header["Client-Id"] == nil || request.Header["Client-Id"][0] == "" {
			helpers.HandleResponse(response, 400, "client id not found", nil)

			return
		}

		next(response, request)
	})
}

func Timestamp(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Header["Timestamp"] == nil || request.Header["Timestamp"][0] == "" {
			helpers.HandleResponse(response, 400, "timestamp not found", nil)

			return
		}

		next(response, request)
	})
}

func AppID(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Header["App-Id"] == nil || request.Header["App-Id"][0] == "" {
			helpers.HandleResponse(response, 400, "app id not found", nil)

			return
		}

		next(response, request)
	})
}

func PrivateKey(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Header["Private-Key"] == nil || request.Header["Private-Key"][0] == "" {
			helpers.HandleResponse(response, 400, "private key not found", nil)

			return
		}

		next(response, request)
	})
}
