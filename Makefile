build:
	docker build -t delegacky .

wire:
	wire ./internal

local:
	docker run -p 5000:80 --env-file .env delegacky