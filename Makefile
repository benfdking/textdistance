pre_bench_branch_name := $(shell git branch --show-current)

.PHONY: benchAgainstMaster
benchAgainstMaster:
	@git diff --exit-code 
	@go test -test.benchmem=true -run=NONE -bench=. ./... > new.txt
	@git co master
	@go test -test.benchmem=true -run=NONE -bench=. ./... > old.txt
	benchcmp old.txt new.txt
	@git co  $(pre_bench_branch_name)
