SHORT_SHA=$(shell git rev-parse --verify --short=8 HEAD)

########## Store Gateway ##########
build-gateway-store: 
	docker build \
		-f src/svc/gateway-store/Dockerfile \
		-t "ghcr.io/briand787b/ncx-gateway-store:$(SHORT_SHA)" \
        -t ghcr.io/briand787b/ncx-gateway-store:latest \
		src/svc

push-gateway-store:
	docker push ghcr.io/briand787b/ncx-gateway-store:$(SHORT_SHA) && \
    docker push ghcr.io/briand787b/ncx-gateway-store:latest

local-gateway-store: build-gateway-store
	docker tag ghcr.io/briand787b/ncx-gateway-store localhost:32000/briand787b/ncx-gateway-store
	docker push localhost:32000/briand787b/ncx-gateway-store

########## Web Gateway ##########
build-gateway-web: 
	docker build \
		-f src/svc/gateway-web/Dockerfile \
		-t "ghcr.io/briand787b/ncx-gateway-web:$(SHORT_SHA)" \
        -t ghcr.io/briand787b/ncx-gateway-web:latest \
		src/svc

push-gateway-web:
	docker push ghcr.io/briand787b/ncx-gateway-web:$(SHORT_SHA) && \
    docker push ghcr.io/briand787b/ncx-gateway-web:latest

local-gateway-web: build-gateway-web
	docker tag ghcr.io/briand787b/ncx-gateway-web localhost:32000/briand787b/ncx-gateway-web
	docker push localhost:32000/briand787b/ncx-gateway-web

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
		-f src/ui/pos/Dockerfile \
		-t "ghcr.io/briand787b/ncx-pos:$(SHORT_SHA)" \
        -t ghcr.io/briand787b/ncx-pos:latest \
		src/ui/pos

push-pos:
	docker push ghcr.io/briand787b/ncx-pos:$(SHORT_SHA) && \
    docker push ghcr.io/briand787b/ncx-pos:latest

local-pos: build-pos
	docker tag ghcr.io/briand787b/ncx-pos localhost:32000/briand787b/ncx-pos
	docker push localhost:32000/briand787b/ncx-pos

########## UTILITIES ##########
js-exec-vol:
	docker-compose run $(shell [[ $(SVC) =~ '-ui' ]] && echo $(SVC) || echo $(SVC)-ui) /bin/bash