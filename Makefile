check:
	go test ./... -race

bench:
	mv bench_results.txt old_results.txt
	go test -tags bench -bench=. $(glide novendor) -benchmem | tee bench_results.txt
	benchcmp old_results.txt bench_results.txt |tee bench_diff.txt
	rm old_results.txt
