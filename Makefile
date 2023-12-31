build:
	@go build -o ./bin/ocha ./ocha.go

run: build
	@./bin/ocha notes ls
