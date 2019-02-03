.PHONY: setup, build_cat, run_cat, build_person, run_person

all: main

main: setup build_cat build_person

setup:
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go
build_cat:
	go build -o cat/client cat/client.go
	go build -o cat/server cat/server.go
build_okashi:
	go build -o okashi/client okashi/client.go
	go build -o okashi/server okashi/server.go
run_cat:
	./cat/server &
	./cat/client
build_person:
	go build -o person/client person/client.go
	go build -o person/server person/server.go
run_person:
	./person/server &
	./person/client
run_okashi:
	./okashi/server &
	./okashi/client

