prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest
	@echo "Install Husky"
	go install github.com/go-courier/husky/cmd/husky@latest && husky init

test:
	go test -v -cover ./...

.PHONY: prepare test
