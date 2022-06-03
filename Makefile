SHORT_SHA=$(shell git rev-parse --verify --short=8 HEAD)

build-ide: 
	docker build \
		-f src/ui/ide/Dockerfile \
		-t "ghcr.io/briand787b/ncx-ide:$(SHORT_SHA)" \
        -t ghcr.io/briand787b/ncx-ide:latest \
		src/ui/ide

push-ide:
	docker push ghcr.io/briand787b/ncx-ide:$(SHORT_SHA) && \
    docker push ghcr.io/briand787b/ncx-ide:latest