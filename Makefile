production:
	go build src/main.go -ldflags "-s -w" -o rrn
	upx rrn

dev:
	go build src/main.go -o dev-rrn