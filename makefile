export GO111MODULE=on
SOLOBOAT_LDFLAGS += -X "soloos/soloboat/version.BuildTS=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
SOLOBOAT_LDFLAGS += -X "soloos/soloboat/version.GitHash=$(shell git rev-parse HEAD)"
# SOLOBOAT_PREFIX += GOTMPDIR=./go.build/tmp GOCACHE=./go.build/cache

all:soloboatd

soloboatd:
	$(SOLOBOAT_PREFIX) go build -i -ldflags '$(SOLOBOAT_LDFLAGS)' -o ./bin/soloboatd ./apps/soloboatd

sidecard:
	$(SOLOBOAT_PREFIX) go build -i -ldflags '$(SOLOBOAT_LDFLAGS)' -o ./bin/sidecard ./apps/sidecard

include ./make/test
include ./make/bench

.PHONY:all soloboatd sidecard test
