# Variables
GO = go
TEST_FLAGS = -v

build:
	$(GO) build -o heroku

test:
	$(GO) test $(TEST_FLAGS) ./...

cover:
	$(GO) test -cover ./...