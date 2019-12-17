.DEFAULT_GOAL: all


ifneq ($(wildcard .env),)
    include .env
    export $(shell sed 's/=.*//' .env)
endif


all: build local

build:
	docker build -t delegacky .

wire:
	wire ./internal

local:
	docker rm szlh-dlg --force
	docker run -p 5000:80 --env-file .env.docker -d --name szlh-dlg delegacky

stop:
	docker stop szlh-dlg

parse:
	docker exec szlh-dlg ./parser

migrate:
	@migrate -path migrations -database "${DB_DRIVER}://${DATABASE_URL}" drop
	@migrate -path migrations -database "${DB_DRIVER}://${DATABASE_URL}" up