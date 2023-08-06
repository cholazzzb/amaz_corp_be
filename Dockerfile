# syntax=docker/dockerfile:1

FROM golang:alpine AS build-stage

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /ac-be

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

COPY --from=build-stage /ac-be /app
COPY --from=build-stage /.env /app
COPY --from=build-stage /migration /app/migration

EXPOSE 8080

USER nonroot:nonroot

CMD ["/app/ac-be"]


