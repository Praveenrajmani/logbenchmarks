
```sh
04:41:54 praveen@piper logbenchmarks → go test -bench=. -benchtime=100x ./...
goos: linux
goarch: amd64
pkg: github.com/miniohq/logbenchmarks
cpu: Intel(R) Core(TM) Ultra 7 165U
BenchmarkJSONSerialization-14          	     100	       807.5 ns/op
BenchmarkJSONDeserialization-14        	     100	      2301 ns/op
BenchmarkJSONZstdCompression-14        	     100	    838872 ns/op
BenchmarkProtobufSerialization-14      	     100	       621.9 ns/op
BenchmarkProtobufDeserialization-14    	     100	       640.5 ns/op
BenchmarkProtobufZstdCompression-14    	     100	    828402 ns/op
PASS

ok  	github.com/miniohq/logbenchmarks	0.185s
04:41:58 praveen@piper logbenchmarks → go test -bench=. -benchtime=1000x ./...
goos: linux
goarch: amd64
pkg: github.com/miniohq/logbenchmarks
cpu: Intel(R) Core(TM) Ultra 7 165U
BenchmarkJSONSerialization-14          	    1000	       420.1 ns/op
BenchmarkJSONDeserialization-14        	    1000	      1818 ns/op
BenchmarkJSONZstdCompression-14        	    1000	    889594 ns/op
BenchmarkProtobufSerialization-14      	    1000	       364.4 ns/op
BenchmarkProtobufDeserialization-14    	    1000	       417.0 ns/op
BenchmarkProtobufZstdCompression-14    	    1000	    786671 ns/op
PASS

ok  	github.com/miniohq/logbenchmarks	1.695s
04:42:06 praveen@piper logbenchmarks → go test -bench=. -benchtime=10000x ./...
goos: linux
goarch: amd64
pkg: github.com/miniohq/logbenchmarks
cpu: Intel(R) Core(TM) Ultra 7 165U
BenchmarkJSONSerialization-14          	   10000	       525.8 ns/op
BenchmarkJSONDeserialization-14        	   10000	      2058 ns/op
BenchmarkJSONZstdCompression-14        	   10000	    890731 ns/op
BenchmarkProtobufSerialization-14      	   10000	       280.0 ns/op
BenchmarkProtobufDeserialization-14    	   10000	       555.9 ns/op
BenchmarkProtobufZstdCompression-14    	   10000	    924503 ns/op
PASS
ok  	github.com/miniohq/logbenchmarks	18.208s

04:43:01 praveen@piper logbenchmarks → go test -bench=. -benchtime=100000x ./...
goos: linux
goarch: amd64
pkg: github.com/miniohq/logbenchmarks
cpu: Intel(R) Core(TM) Ultra 7 165U
BenchmarkJSONSerialization-14          	  100000	       434.4 ns/op
BenchmarkJSONDeserialization-14        	  100000	      1892 ns/op
BenchmarkJSONZstdCompression-14        	  100000	    888789 ns/op
BenchmarkProtobufSerialization-14      	  100000	       278.9 ns/op
BenchmarkProtobufDeserialization-14    	  100000	       399.5 ns/op
BenchmarkProtobufZstdCompression-14    	  100000	    926719 ns/op
PASS
ok  	github.com/miniohq/logbenchmarks	181.871s
```
