CMD = ./pyro.go
RUN = go run -race ${CMD}
LDFLAGS = -ldflags="-s -w"

all: run run-test install

run:
	${RUN}; true # don't care about pyro's exit code

run-test:
	${RUN} test test/test.pyro

# build for specific OS target
build-%:
	GOOS=$* GOARCH=amd64 go build ${LDFLAGS} -o pyro-$* ${CMD}


# build for all OS targets, useful for releases
build: build-linux build-darwin build-windows build-openbsd


# install on host system
install:
	go install ${LDFLAGS} ${CMD}
	ls -l `which pyro`


# pre-commit hook
precommit:
	go vet ./cmd ./pkg/pyro
	go fmt ./cmd ./pkg/pyro


# clean any generated files
clean:
	rm -rvf pyro pyro-*
