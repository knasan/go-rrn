production:
	go build -ldflags "-s -w" -o rrn
	upx rrn

dev:
	go build -o dev-rrn