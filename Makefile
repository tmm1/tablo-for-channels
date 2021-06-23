docker-build:
	docker buildx build --push --platform linux/amd64,linux/arm64,linux/arm/v7 . -t tmm1/tablo-for-channels:latest

.PHONY: docker-build
