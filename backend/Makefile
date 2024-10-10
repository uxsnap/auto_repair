.PHONY: docker
	docker:
		open -a docker

.PHONY: run
run:
	nodemon --watch './internal/**/*.go' --signal SIGTERM --exec 'go' run ./cmd/main.go