.PHONY: main prac

main:
	@go build -o bin/$@ src/main.go src/models.go

prac:
	@go build -o bin/$@ src/prac.go

clean:
	@rm bin/*
