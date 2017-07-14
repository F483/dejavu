PROFILE_ARGS := -cpuprofile cpu.prof -memprofile mem.prof

PROJECT_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
PROJECT_NAME := $(notdir $(patsubst %/,%,$(dir $(PROJECT_PATH))))

doc:
	godoc .

clean:
	@rm -f *.prof
	@rm -f *.test

install:
	go install
	cd $(PROJECT_NAME) && go install

test:
	go test

test_single:
	go test -run $(TESTNAME)

coverage_annotate:
	gocov test | gocov annotate -

coverage_report:
	gocov test | gocov report

profile_test:
	go test $(PROFILE_ARGS) .

profile_test_single:
	go test $(PROFILE_ARGS) -run $(TESTNAME)

profile_results_cpu:
	go tool pprof cpu.prof

profile_results_mem:
	go tool pprof http://localhost:6060/debug/pprof/heap
