APP=sph-backend-coding-challenge
APP_EXECUTABLE="./out/$(APP)"
ALL_PACKAGES=$(shell go list ./... | grep -v "mocks" | grep -v "vendor" | grep -v "cmd")
NON_VENDOR_PACKAGES=$(shell go list ./... | grep -v "vendor")
UNIT_TEST_PACKAGES=$(shell go list ./... | grep -v "it")
# optionally pass this from command line
MIGRATION_ROOT?=./db/migration
# optionally pass this from command line
DB?=sph_development
# optionally pass this from command line
PORT?=8080
ENV_LOCAL_TEST=\
				DB_HOST=localhost \
				DB_PORT=5432 \
				DB_NAME=$(DB) \
				DB_USER=postgres \
				DB_PASS=mysecretpassword \
				DB_MAX_IDLE_CONN=10 \
				DB_MAX_OPEN_CONN=200 \
				DB_CONN_MAX_LIFETIME=30m \
				ENV=local\
				PORT=8080

setup:
		go get -u github.com/golang/dep/cmd/dep
		go get -u github.com/golang/lint/golint

build-deps:
		dep ensure

compile:
		mkdir -p out/
		go build -o $(APP_EXECUTABLE)

build: build-deps compile fmt vet

install:
		go install ./...

fmt:
		go fmt $(NON_VENDOR_PACKAGES)

vet:
		go vet $(NON_VENDOR_PACKAGES)

docker.start:
		docker-compose up -d;

docker.stop:
		docker-compose down;

docker.restart: docker.stop docker.start

docker.sph.app.stop: 
		docker stop sph-backend-coding-challenge_sph-server_1

test.unit:
		go test $(UNIT_TEST_PACKAGES) --cover -race -v -count 1

test.update: docker.sph.app.stop 
		$(ENV_LOCAL_TEST) \
		go test -tags=golden ./it -v -count=1 -update

test.integration: docker.sph.app.stop
		$(ENV_LOCAL_TEST) \
		go test -tags=integration ./it -v -count=1

migrate.up:
		migrate -path $(MIGRATION_ROOT) -database "postgres://postgres:mysecretpassword@localhost:5432/$(DB)?sslmode=disable" up

migrate.down:
		migrate -path $(MIGRATION_ROOT) -database "postgres://postgres:mysecretpassword@localhost:5432/$(DB)?sslmode=disable" down