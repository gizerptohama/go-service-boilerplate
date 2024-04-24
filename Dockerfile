# Build the application from source
FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build -o ./boilerplate-app ./main.go

# Deploy the application binary into a lean image
FROM alpine:3.19.1 AS build-release-stage

WORKDIR /

COPY --from=build-stage ./app/boilerplate-app ./boilerplate-app

EXPOSE 8080

ENTRYPOINT ["/boilerplate-app"]
