NAME := gsql

ifeq ($(OS),Windows_NT)
	EXT = .exe
endif

.DEFAULT_GOAL: $(NAME)

$(NAME):
	go build -o ${NAME}${EXT} cmd/gsql/main.go

clean:
	rm -f ${NAME}

setup:
	go get "github.com/Masterminds/glide"

test:

lint:


.PHONY: clean setup test lint

