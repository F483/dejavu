
setup:
	go get github.com/axw/gocov/gocov
	go get golang.org/x/tools/cmd/godoc


install:
	go install


doc:
	godoc github.com/f483/dejavu


test:
	go test -v


lint:
	# TODO golint


benchmark:
	# TODO benchmark


example:
	go run example.go


coverage:
	gocov test | gocov annotate -
