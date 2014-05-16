default: build

assets: assets/bindata.go

# go get github.com/jteeuwen/go-bindata/...
assets/bindata.go: assets/*.css assets/*.js
	@cd assets && go-bindata -pkg=assets -ignore=bindata.go .

build: rationl.pb.go

# go get github.com/benbjohnson/ego/...
ego:
	@ego .

rationl.pb.go: rationl.proto
	protoc --gogo_out=. -I=.:../../../ rationl.proto

run: rationl.pb.go assets ego
	@mkdir -p tmp
	go run ./cmd/rationl/main.go -data-dir tmp -client-id="$(RATIONL_CLIENT_ID)" -secret="$(RATIONL_SECRET)"

.PHONY: assets
