#FROM golang:1.12-alpine AS builder-env
FROM golang:1.12 AS builder-env

WORKDIR /app

#RUN apk add --no-cache git

ADD ./cmd ./cmd
ADD ./internal ./internal
ADD ./go.mod .
ADD ./go.sum .

RUN go mod vendor

ENV CGO_ENABLED 0
ENV GOBIN /usr/local/bin/

RUN go get github.com/google/wire/cmd/wire
RUN wire ./internal

RUN go install -a ./...



FROM alpine:latest

COPY --from=builder-env /usr/local/bin/web .

CMD ["./web"]