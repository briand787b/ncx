module github.com/briand787b/ncx/src/svc/gateway-store

go 1.18

require (
	github.com/briand787b/ncx/src/svc/pkg v0.0.0
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/render v1.0.2
	github.com/google/uuid v1.3.0
	github.com/pkg/errors v0.9.1
)

require (
	github.com/ajg/form v1.5.1 // indirect
	golang.org/x/net v0.0.0-20220822230855-b0a4917ee28c // indirect
)

replace github.com/briand787b/ncx/src/svc/pkg v0.0.0 => ../pkg
