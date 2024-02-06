CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
MOCKGEN=${BINDIR}/mockgen_${GOVER}
PACKAGE=github.com/mystpen/TG-bot/cmd/bot

build: bindir
	go build -o ${BINDIR}/bot ${PACKAGE}

test:
	go test ./...
	
run:
	go run ${PACKAGE}

generate:
	${MOCKGEN} -source=internal/model/messages/incoming_msg.go -destination=internal/mocks/messages/messages_mock.go

bindir:
	mkdir -p ${BINDIR}