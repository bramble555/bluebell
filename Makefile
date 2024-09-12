
# 指令
.PHONY: all build run gotool clean help tidy

BINARY=bluebell

all: gotool build

# build:
# 	CGO_ENABLED=0 GOOS=Windows GOARCH=amd64 go build -o ${BINARY}
build:
	go build -o ${BINARY}.exe & ${BINARY}

run:
	go run main.go
tidy:
	go mod tidy

gotool:
	go fmt ./
	go vet ./

# clean:
# 	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	# @echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"