# syntax=docker/dockerfile:1

FROM golang:1.16-buster AS build

WORKDIR /go/src/app

COPY backend/go.mod ./
COPY backend/go.sum ./

RUN go mod download

COPY /backend/*.go ./

RUN go build -o /uni-tasks

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY backend/server.crt ./
COPY backend/server.key ./

COPY --from=build /uni-tasks /uni-tasks

EXPOSE 8080
EXPOSE 443

USER nonroot:nonroot

ENTRYPOINT ["/uni-tasks"]