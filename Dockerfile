FROM golang:1.12-alpine AS builder-env

RUN apk --no-cache add ca-certificates git

ENV CGO_ENABLED 0
ENV GOBIN /usr/local/bin/

WORKDIR /app

ADD ./server/cmd ./cmd
ADD ./server/internal ./internal
ADD ./server/go.mod .
ADD ./server/go.sum .

RUN go mod vendor

RUN go get github.com/google/wire/cmd/wire
RUN wire ./internal

RUN go install -a ./...


FROM alpine:latest

RUN apk add --update tzdata
ENV TZ=Europe/Warsaw
RUN rm -rf /var/cache/apk/*

COPY --from=builder-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder-env /usr/local/bin/web .
COPY --from=builder-env /usr/local/bin/parser .

CMD ["./web"]