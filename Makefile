VERSION = `cat VERSION`
echo:
	echo $(VERSION)
build:
	GOOS=linux go build -o main cmd/main.go && docker build -t asia.gcr.io/instant-matter-785/street_name:$(VERSION) . && rm main
publish:
	gcloud docker push asia.gcr.io/instant-matter-785/street_name:$(VERSION)
test:
	go test
run:
	go run cmd/main.go -http.addr :3456
