NAME := gsql

ifeq ($(OS),Windows_NT)
	EXT = .exe
endif

.DEFAULT_GOAL: $(NAME)

$(NAME):
	go build -o ${NAME}${EXT} cmd/gsql/main.go

.PHONY: clean
clean:
	rm -f ${NAME}

.PHONY: setup
setup:
	go get "github.com/mitchellh/cli"
	go get "github.com/lib/pq"
	go get "gopkg.in/yaml.v2"

.PHONY: test
test:

.PHONY: lint
lint:



