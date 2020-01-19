.PHONY: main prac

main:
	@go build -o bin/$@ src/main.go src/models.go src/handlers.go src/consts.go

prac:
	@go build -o bin/$@ src/prac.go

clean:
	@rm bin/*
