# pedal-planner-api-v1

## Description

Pedal Planner: API Service V1 is the backend service for the Pedal Planner application.

## Getting Started

[Golang](https://golang.org/) 1.15 or higher will be needed to run this service

```
brew install go
```

## Build

A `Makefile` is provided in the project root to facilitate running this project

### start application

```
make up
```

### stop application

```
make down
```

### update dependencies

```
make vendor
```

## Documentation

Documentation of API routes is done using [Swagger UI](https://swagger.io).

The documentation is located in the [Pedal Planner API v1 Spec repo.](https://github.com/zack-jack/pedal-planner-api-v1-spec)

### update docs

```
make docs
```

To view the API docs, run `make up` and navigate to the `/docs` route.
