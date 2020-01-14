.PHONY: main prac

main:
	@go build -o bin/$@ src/main.go

prac:
	@go build -o bin/$@ src/prac.go

