# syntax=docker/dockerfile:1

FROM golang:1.16-buster AS build

WORKDIR /go/src/app

COPY backend/go.mod ./
COPY backend/go.sum ./

RUN go mod download

ADD backend/. /go/src/app/
COPY /backend/*.go ./

RUN go build -o /uni-tasks

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY backend/*.pgsql ./
COPY backend/*.sql ./
COPY backend/config.sample.json ./

COPY --from=build /uni-tasks /uni-tasks

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/uni-tasks"]