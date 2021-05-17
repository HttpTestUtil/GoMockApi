release: mkdir
	go build -o release/mock-api

mkdir:
	rm -rf ./release
	mkdir release
