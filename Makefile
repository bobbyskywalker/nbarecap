build:
	go build -o nbarecap main.go
run:
	./nbarecap games
clean:
	rm -rf nbarecap
cleanall:
	rm -rf nbarecap && rm -rf tea.log
test:
	go test -v ./internal/nba && echo "\n" && go test -v ./pkg/nba_api/mappers

