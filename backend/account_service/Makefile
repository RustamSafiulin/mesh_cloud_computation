executable := account_service

all: $(executable)

clean:
	@rm -rf bin/* || true
	@rm -rf bin || true

$(executable):
	@echo Building $@
	go build -o bin/$@ cmd/$@/main.go 

build:
	@echo Building $(executable)
	go build -o bin/$(executable) cmd/$(executable)/main.go 

.PHONY: $(executable) clean build
