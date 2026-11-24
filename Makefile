.PHONY: build run

EXT :=
ifeq ($(OS),Windows_NT)
    EXT := .exe
endif

BIN_DIR := ./bin

build:
	@clear
	@mkdir -p $(BIN_DIR)
	@go build -o ./bin/main$(EXT) ./Cmd/Main

run: build
	@./bin/main$(EXT)
