# make all - build all things
all:
	-gotags -R . > tags
	go build
	go test -c
	rm -f ippx.test

# make test - run tests
test:
	go test

# make cover - open coverage window in browser
cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
