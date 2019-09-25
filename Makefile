
build docker : *.go
	GOARCH="amd64" GOOS="linux" go build -o webapp && \
	docker-compose up
