FROM golang:1.12-alpine AS goland_builder

RUN apk --no-cache add git

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

# Build the React application
FROM node:alpine AS node_builder

COPY ./client ./
RUN npm install
RUN npm run build


# Final container
FROM alpine:latest

RUN apk add --update tzdata
ENV TZ=Europe/Warsaw
RUN rm -rf /var/cache/apk/*

RUN apk --no-cache add ca-certificates

COPY --from=goland_builder /usr/local/bin/api .
COPY --from=goland_builder /usr/local/bin/parser .
RUN chmod +x ./api

COPY --from=node_builder /build ./web

EXPOSE 8080

CMD ["./api"]