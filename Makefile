GIT_TAG=`git describe --tags --abbrev=0`
COMMIT=`git rev-parse --short HEAD`
PACKR=$(which packr2)
BUILD_DATE=$(shell date +%FT%T%z)

LDFLAGS="-X=github.com/lukasjarosch/godin/internal.Version=$(GIT_TAG) \
		-X=github.com/lukasjarosch/godin/internal.Commit=$(COMMIT) \
		-X=github.com/lukasjarosch/godin/internal.Build=${BUILD_DATE}"

build:
	@GO11MODULE=on packr2 --legacy build -o ./bin/godin -ldflags ${LDFLAGS}

install:
	@GO11MODULE=on packr2 --legacy install -ldflags ${LDFLAGS}
