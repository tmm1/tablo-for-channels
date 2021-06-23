docker-build:
	docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 . -t tmm1/tablo-for-channels:latest

docker-upload: docker-build
	docker push tmm1/tablo-for-channels:latest

.PHONY: docker-build docker-upload
