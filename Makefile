PROFILEARGS := -cpuprofile cpu.prof -memprofile mem.prof

doc:
	godoc .

clean:
	rm *.prof
	rm *.test

install:
	go install
	cd dejavu && go install

test:
	go test

test_single:
	go test -run $(TESTNAME)

coverage_annotate:
	gocov test | gocov annotate -

coverage_report:
	gocov test | gocov report

profile_test:
	go test $(PROFILEARGS) .

profile_test_single:
	go test $(PROFILEARGS) -run $(TESTNAME)

profile_results_cpu:
	go tool pprof cpu.prof

profile_results_mem:
	go tool pprof http://localhost:6060/debug/pprof/heap
