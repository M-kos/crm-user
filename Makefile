generate-openapi:
	openapi-generator generate -g go-server -i ./openapi/openapi.yaml -o internal/api -c ./openapi/config.yaml

build:
	go build -o ./bin/user ./cmd/main.go

build-with-generate: generate-openapi build

lint:
	golangci-lint run ./... -v

clean-bin:
	rm -rf ./bin

tidy:
	go mod tidy -v

# окружение: api, db, cache
dev-up:
	echo "dev-up"

# удаляется окружение и volumes
dev-down:
	echo "dev-down"

# curl проверка
dev-api-test:
	echo "dev-api-test"

# заходит в базу - psql -U postgres -d dev
dev-db:
	echo "dev-db"

# запускает локально api
local-api-up:
	CRM_ENV_FILE=".env" go run ./cmd/main.go

# https://github.com/githubnemo/CompileDaemon
local-api-up-hot:
	CompileDaemon -exclude-dir=.git -exclude-dir=bin -exclude-dir=openapi -build='make build' -command='./bin/user' -color=true

docker-build:
	docker build -t crm-user:latest -f ./Dockerfile .

docker-run:
	docker run --name user-service -d -p 8080:8080 -e CRM_PORT=8080 crm-user:latest

compose-up:
	docker compose -f docker-compose.yaml --project-directory=./ up -d --build

compose-down:
	docker compose -f docker-compose.yaml --project-directory=./ down


.PHONY: generate-openapi build build-with-generate lint clean-bin tidy dev-up dev-down dev-api-test dev-db local-api-up local-api-up-hot docker-build docker-run compose-up compose-down