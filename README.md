# Syncbench

Trivial synthentic micro-benchmark of Go synchronization primitives.

Results on old 2009 iMac:

    dev@saturn:~/hashrocket/go/src/github.com/jackc/syncbench% go test -bench=.
    testing: warning: no tests to run
    BenchmarkNoContention-4           2000000000           0.63 ns/op
    BenchmarkMutexNoContention-4      50000000          24.4 ns/op
    BenchmarkMutexContention-4        20000000          69.6 ns/op
    BenchmarkAtomicNoContention-4     200000000          8.52 ns/op
    BenchmarkAtomicContention-4       50000000          25.4 ns/op
    BenchmarkChannelSelect-4           2000000         734 ns/op
    BenchmarkChannelRange-4            3000000         401 ns/op
    PASS
    ok    github.com/jackc/syncbench  11.757s
