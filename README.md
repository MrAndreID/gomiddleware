# MrAndreID / Go Middleware

[![Go Reference](https://pkg.go.dev/badge/github.com/MrAndreID/gomiddleware.svg)](https://pkg.go.dev/github.com/MrAndreID/gomiddleware) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The `MrAndreID/GoMiddleware` package is a collection of functions in the go language.

---

## Table of Contents

* [Install](#install)
* [Usage](#usage)
* [Versioning](#versioning)
* [Authors](#authors)
* [License](#license)
* [Official Documentation for Go Language](#official-documentation-for-go-language)
* [More](#more)

---

## Install

To use The `MrAndreID/GoMiddleware` package, you must follow the steps below:

```sh
go get -u github.com/MrAndreID/gomiddleware
```

## Usage

To use The `MrAndreID/GoMiddleware` package, you must combine it with The `gorilla/mux` package.

### Log

```go
router := mux.NewRouter().StrictSlash(true)

router.Handle("/api/v1/", gomiddleware.Log(Documentation)).Methods("GET")
```

### Acceptable

```go
router := mux.NewRouter().StrictSlash(true)

router.Handle("/api/v1/", gomiddleware.Acceptable(Documentation)).Methods("GET")
```

### Client ID

```go
router := mux.NewRouter().StrictSlash(true)

router.Handle("/api/v1/", gomiddleware.ClientID(Documentation)).Methods("GET")
```

### Timestamp

```go
router := mux.NewRouter().StrictSlash(true)

router.Handle("/api/v1/", gomiddleware.Timestamp(Documentation)).Methods("GET")
```

### App ID

```go
router := mux.NewRouter().StrictSlash(true)

router.Handle("/api/v1/", gomiddleware.AppID(Documentation)).Methods("GET")
```

## Versioning

I use [SemVer](https://semver.org/) for versioning. For the versions available, see the tags on this repository. 

## Authors

**Andrea Adam** - [MrAndreID](https://github.com/MrAndreID/)

## License

MIT licensed. See the LICENSE file for details.

## Official Documentation for Go Language

Documentation for Go Language can be found on the [Go Language website](https://golang.org/doc/).

## More

Documentation can be found [on https://go.dev/](https://pkg.go.dev/github.com/MrAndreID/gomiddleware).
