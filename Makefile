docker-build:
	docker build . -t tmm1/tablo-for-channels:latest

docker-upload: docker-build
	docker push tmm1/tablo-for-channels:latest

.PHONY: docker-build docker-upload
