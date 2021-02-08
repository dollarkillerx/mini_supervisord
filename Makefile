build:
	CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -o mini_suppervisord cmd/main.go

upx:
	upx mini_suppervisord