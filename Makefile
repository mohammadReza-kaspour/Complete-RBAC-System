.PHONY: build run dirmap

EXT :=
ifeq ($(OS),Windows_NT)
    EXT := .exe
endif

BIN_DIR := ./bin
dirmap: 
	@powershell -NoProfile -Command "Get-ChildItem -Recurse -Name > dirmap.log"

build: 	dirmap
		@clear
		@mkdir -p $(BIN_DIR)
		@go build -o ./bin/main$(EXT) ./cmd/server


run: 	build
		@./bin/main$(EXT)

