SHORT_SHA=$(shell git rev-parse --verify --short=8 HEAD)

########## IDE ##########
build-ide: 
	docker build \
		-f src/ui/ide/Dockerfile \
		-t "ghcr.io/briand787b/ncx-ide:$(SHORT_SHA)" \
        -t ghcr.io/briand787b/ncx-ide:latest \
		src/ui/ide

push-ide:
	docker push ghcr.io/briand787b/ncx-ide:$(SHORT_SHA) && \
    docker push ghcr.io/briand787b/ncx-ide:latest

local-ide: build-ide
	docker tag ghcr.io/briand787b/ncx-ide localhost:32000/briand787b/ncx-ide
	docker push localhost:32000/briand787b/ncx-ide

########## POS ##########
build-pos: 
	docker build \
		-f src/ui/ide/Dockerfile \
		-t "ghcr.io/briand787b/ncx-pos:$(SHORT_SHA)" \
        -t ghcr.io/briand787b/ncx-pos:latest \
		src/ui/pos

push-pos:
	docker push ghcr.io/briand787b/ncx-pos:$(SHORT_SHA) && \
    docker push ghcr.io/briand787b/ncx-pos:latest

local-pos: build-pos
	docker tag ghcr.io/briand787b/ncx-pos localhost:32000/briand787b/ncx-pos
	docker push localhost:32000/briand787b/ncx-pos