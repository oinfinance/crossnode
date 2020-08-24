all: install

install: go.sum
		go install -tags "$(build_tags)" ./cmd/oind

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		@go mod verify
