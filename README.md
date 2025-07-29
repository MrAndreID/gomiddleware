# MrAndreID / Go Middleware

[![Go Reference](https://pkg.go.dev/badge/github.com/MrAndreID/gomiddleware.svg)](https://pkg.go.dev/github.com/MrAndreID/gomiddleware) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The `MrAndreID/GoMiddleware` package is a collection of functions in the go language.

## Table of Contents

* [Installation](#installation)
* [Usage](#usage)
* [Versioning](#versioning)
* [Authors](#authors)
* [Contributing](#contributing)
* [Official Documentation for Go Language](#official-documentation-for-go-language)
* [License](#license)
* [More](#more)

## Installation

To use The `MrAndreID/GoMiddleware` package, you must follow the steps below:

```go
go get -u github.com/MrAndreID/gomiddleware
```

## Usage

To use The `MrAndreID/GoMiddleware` package, you must combine it with The `gorilla/mux` package or The `Echo Framework`.

### Log

```go
router := mux.NewRouter().StrictSlash(true)

router.Handle("/api/v1/", gomiddleware.Log(Documentation)).Methods("GET")
```

### Acceptable

```go
router := mux.NewRouter().StrictSlash(true)

router.Handle("/api/v1/", gomiddleware.Acceptable(Documentation)).Methods("POST")
```

### Client ID

```go
router := mux.NewRouter().StrictSlash(true)

router.Handle("/api/v1/", gomiddleware.ClientID(Documentation)).Methods("POST")
```

### Timestamp

```go
router := mux.NewRouter().StrictSlash(true)

router.Handle("/api/v1/", gomiddleware.Timestamp(Documentation)).Methods("POST")
```

### App ID

```go
router := mux.NewRouter().StrictSlash(true)

router.Handle("/api/v1/", gomiddleware.AppID(Documentation)).Methods("POST")
```

### Private Key

```go
router := mux.NewRouter().StrictSlash(true)

router.Handle("/api/v1/", gomiddleware.PrivateKey(Documentation)).Methods("POST")
```

### Content Type

```go
router := mux.NewRouter().StrictSlash(true)

router.Handle("/api/v1/", gomiddleware.ContentType(Documentation)).Methods("POST")
```

## Versioning

I use [Semantic Versioning](https://semver.org/). For the versions available, see the tags on this repository. 

## Authors

**Andrea Adam** - [MrAndreID](https://github.com/MrAndreID/)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.

## Official Documentation for Go Language

Documentation for Go Language can be found on the [Go Package website](https://pkg.go.dev/).

## License

The `MrAndreID/GoMiddleware` is released under the [MIT License](https://opensource.org/licenses/MIT). See the `LICENSE` file for more information.

## More

Documentation can be found [on https://go.dev/](https://pkg.go.dev/github.com/MrAndreID/gomiddleware).
