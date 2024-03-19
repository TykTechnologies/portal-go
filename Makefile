COPYRIGHT_HOLDER := "Tyk Technologies"
COPYRIGHT_YEARS := "2024"

.PHONY: update-deps
update-deps:
	go get -u ./...

.PHONY: test
test:
	go test -race -v ./...

.PHONY: dep
dep:
	go mod download

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: addlicense
addlicense:
	go install github.com/google/addlicense@latest

.PHONY: copyright
copyright: #addlicense
	addlicense -c ${COPYRIGHT_HOLDER} -y ${COPYRIGHT_YEARS} -l mpl -s=only .

.PHONY: check-license
check-license: addlicense
	addlicense -check .
