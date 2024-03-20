all:
	-gotags -R . > tags
	go build
	go test -c
	rm -f ippx.test

test:
	go test
