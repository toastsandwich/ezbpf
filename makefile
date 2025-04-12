.PHONY: grammar build
grammar:
	@antlr -o grammar -package grammar -Dlanguage=Go -visitor -no-listener ezbpf.g4

build:
	@go build -o build/ezbpf main.go

run:
	@go run main.go
