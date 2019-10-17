cov=coverage.out
covhtml=coverage.html

check:
	go test ./... -race -coverprofile=$(cov)

coverage: check
	go tool cover -html=$(cov) -o=$(covhtml)
	xdg-open coverage.html

bench:
	mv bench_results.txt old_results.txt
	GOMAXPROCS=4 go test -bench=. ./... -benchmem | tee bench_results.txt
	benchcmp old_results.txt bench_results.txt |tee bench_diff.txt
	rm old_results.txt
