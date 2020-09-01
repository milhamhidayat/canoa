TEST_OPTS := -covermode=atomic $(TEST_OPTS)

MYSQL_USER ?= soccer
MYSQL_PASSWORD ?= soccer-pass
MYSQL_ADDRESS ?= 127.0.0.1:3306
MYSQL_DATABASE ?= soccer

.PHONY: vendor
vendor: go.mod go.sum
	@GO111MODULE=on go get ./...

.PHONY: unittest
unittest: vendor
	go test -short $(TEST_OPTS) ./...

.PHONY: mysql-up
mysql-up:
	@docker-compose up -d mysql

# Database Migration
.PHONY: migrate-prepare
migrate-prepare:
	@GO111MODULE=off go get -tags 'mysql' -u github.com/golang-migrate/migrate/cmd/migrate

.PHONY: migrate-up
migrate-up:
	@migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS))/$(MYSQL_DATABASE)" \
	-path=db/migrations up

.PHONY: migrate-down
migrate-down:
	@migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS))/$(MYSQL_DATABASE)" \
	-path=db/migrations down

.PHONY: migrate-drop
migrate-drop:
	@migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS))/$(MYSQL_DATABASE)" \
	-path=db/migrations drop

.PHONY: hello
hello:
	@echo "hello"
