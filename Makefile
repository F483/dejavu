# TODO add documentation generation


setup:
	go get github.com/axw/gocov/gocov
	go get golang.org/x/tools/cmd/godoc


install:
	go install


doc:
	godoc github.com/f483/dejavu


test:
	go test -v


coverage:
	gocov test | gocov annotate -
