.PHONY: check test build

all: check test build

update:
	@echo "更新"

test:
	@echo "测试"

build: check deps
	@echo "编译"
	GOOS=linux GOARCH=amd64 go build -o web-server main.go
deps:
	@echo "下载依赖"

run: check deps build
	@echo "运行"
	./web-server

