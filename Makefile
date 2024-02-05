PACKAGE=github.com/mystpen/TG-bot/cmd/bot

test:
	go test ./...
	
run:
	go run ${PACKAGE}