release: mkdir
	go release -o build/start

mkdir:
	rm -rf ./release
	mkdir release
	cp -r api release/api
	cp -r config release/config
	cp README.md release/README.md
