GIT_TAG=`git describe --tags --abbrev=0`
COMMIT=`git rev-parse --short HEAD`
PACKR=$(which packr2)
BUILD_DATE=$(shell date +%FT%T%z)

LDFLAGS=-ldflags "-X=main.Version=$(GIT_TAG) -X=main.GitCommit=$(COMMIT) -X=main.BuildDate=${BUILD_DATE}"

build:
	@GO11MODULE=on packr2 --legacy build -o ./bin/godin ${LDFLAGS}

install:
	@GO11MODULE=on packr2 --legacy install ${LDFLAGS}
