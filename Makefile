.PHONY: main prac worker

main:
	@go build -o bin/$@ src/main.go src/models.go src/handlers.go src/consts.go

worker:
	@go build -o bin/$@ src/worker.go src/models.go

prac:
	@go build -o bin/$@ src/prac.go src/models.go

