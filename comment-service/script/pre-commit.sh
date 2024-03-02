#!/bin/sh

echo "Start lint code"
if !(golangci-lint.exe --version); then
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
fi
# 目录名后跟上...表示对该目录进行递归查找
if !(golangci-lint.exe run --enable cyclop --config ./lint.yaml ./...); then
    echo "Lint fail!"
    exit 1
fi

echo "Lint success"

exit 0