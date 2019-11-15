#FROM golang:1.12-alpine AS builder-env
FROM golang:1.12 AS builder-env

ENV CGO_ENABLED 0
ENV GOBIN /usr/local/bin/

WORKDIR /app

ADD ./cmd ./cmd
ADD ./internal ./internal
ADD ./go.mod .
ADD ./go.sum .

RUN go mod vendor

RUN go get github.com/google/wire/cmd/wire
RUN wire ./internal

RUN go install -a ./...


FROM alpine:latest

COPY --from=builder-env /usr/local/bin/web .
COPY --from=builder-env /usr/local/bin/parser .

CMD ["./web"]