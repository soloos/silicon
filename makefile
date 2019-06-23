SILICON_LDFLAGS += -X "soloos/silicon/version.BuildTS=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
SILICON_LDFLAGS += -X "soloos/silicon/version.GitHash=$(shell git rev-parse HEAD)"
# SILICON_PREFIX += GOTMPDIR=./go.build/tmp GOCACHE=./go.build/cache

all:silicond

silicond:
	$(SILICON_PREFIX) go build -i -ldflags '$(SILICON_LDFLAGS)' -o ./bin/silicond ./silicond

include ./make/test
include ./make/bench

.PHONY:all silicond test
