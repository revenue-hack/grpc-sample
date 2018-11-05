.PHONY: setup, build_cat,run_cat

setup:
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go
build_cat:
	go build -o cat/client cat/client.go
	go build -o cat/server cat/server.go
run_cat:
	./cat/server &
	./cat/client

