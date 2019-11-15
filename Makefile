all: build local

build:
	docker build -t delegacky .

wire:
	wire ./internal

local:
	docker rm szlh-dlg --force
	docker run -p 5000:80 --env-file .env -d --name szlh-dlg delegacky

stop:
	docker stop szlh-dlg

parse:
	docker exec szlh-dlg ./parser