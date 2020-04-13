pre_bench_branch_name := $(shell git branch --show-current)

.PHONY: benchAgainstMaster
benchAgainstMaster:
	@git diff --exit-code 
	@go test -test.benchmem=true -run=NONE -bench=. ./... > new.txt
	@git checkout master
	@go test -test.benchmem=true -run=NONE -bench=. ./... > old.txt
	benchcmp old.txt new.txt
	@git checkout  $(pre_bench_branch_name)

.PHONY: fmt
fmt:
	@gofmt -s -w .

.PHONY: test
test:
	@go test --race .