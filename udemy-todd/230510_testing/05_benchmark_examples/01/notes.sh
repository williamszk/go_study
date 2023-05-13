
# to run the benchmarks
cd mystr
go test -bench .
cd -

# goos: linux
# goarch: amd64
# pkg: proj/mystr
# cpu: Intel(R) Core(TM) i5-5300U CPU @ 2.30GHz
# BenchmarkCat-4             14410             82497 ns/op
# BenchmarkJoin-4           484654              3009 ns/op
# PASS
# ok      proj/mystr      4.034s

# it is 4 cores
# Join is much faster...

# ==============================================================================
# to run tests
go test ./...

# ==============================================================================
# to show documentation
godoc -http=:8080
