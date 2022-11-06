run:
	go build -o orderbook-app cmd/app/main.go && ./orderbook-app

run-loadtest:
	go build -o loadtest cmd/loadtest/main.go && ./loadtest