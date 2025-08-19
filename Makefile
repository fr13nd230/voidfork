# Define build command
build:
	@go build -o build/voidfork cmd/*.go

# Define run command
run:
	@go run cmd/*.go

# Define test command
test:
	@go test ./... -v -cover
