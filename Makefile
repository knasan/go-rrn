$(shell mkdir -p build)

prod:
	go build -o build/rrn  -ldflags "-s -w" src/*.go
	upx build/rrn

dev:
	go build -o build/dev-rrn --race src/*.go

run:
	go run --race src/*.go
