FROM golang:1.20-bullseye AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .
COPY internal/ internal/

RUN go build -o /lambda main.go

## Deploy
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /lambda /lambda
COPY assets/ .

USER nonroot:nonroot

ENTRYPOINT ["/lambda"]