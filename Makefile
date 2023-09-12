PKG := github.com/TheRebelOfBabylon/protoc-gen-jsonschema

GOBUILD := go build -v
GOINSTALL := go install -v

# ============
# INSTALLATION
# ============

build:
	$(GOBUILD) -tags="${tags}" -o ./build/${bin} $(PKG)/cmd/protoc-gen-jsonschema

install:
	$(GOINSTALL) -tags="${tags}" $(PKG)/cmd/protoc-gen-jsonschema