
```sh
05:32:53 praveen@piper logbenchmarks ±|main ✗|→ go test -bench=. -benchtime=100x ./...
goos: linux
goarch: amd64
pkg: github.com/miniohq/logbenchmarks
cpu: Intel(R) Core(TM) Ultra 7 165U
BenchmarkJSONSerialization-14             	     100	       817.4 ns/op
BenchmarkJSONDeserialization-14           	     100	      2299 ns/op
BenchmarkJSONZstdCompression-14           	     100	    831558 ns/op
BenchmarkProtobufSerialization-14         	     100	       359.7 ns/op
BenchmarkProtobufDeserialization-14       	     100	      1335 ns/op
BenchmarkProtobufZstdCompression-14       	     100	    850407 ns/op
BenchmarkMessagePackSerialization-14      	     100	       899.9 ns/op
BenchmarkMessagePackZstdCompression-14    	     100	    867546 ns/op
PASS
ok  	github.com/miniohq/logbenchmarks	0.279s

05:33:03 praveen@piper logbenchmarks ±|main ✗|→ go test -bench=. -benchtime=1000x ./...
goos: linux
goarch: amd64
pkg: github.com/miniohq/logbenchmarks
cpu: Intel(R) Core(TM) Ultra 7 165U
BenchmarkJSONSerialization-14             	    1000	       552.3 ns/op
BenchmarkJSONDeserialization-14           	    1000	      2298 ns/op
BenchmarkJSONZstdCompression-14           	    1000	    881865 ns/op
BenchmarkProtobufSerialization-14         	    1000	       401.5 ns/op
BenchmarkProtobufDeserialization-14       	    1000	       440.1 ns/op
BenchmarkProtobufZstdCompression-14       	    1000	    808038 ns/op
BenchmarkMessagePackSerialization-14      	    1000	       625.7 ns/op
BenchmarkMessagePackZstdCompression-14    	    1000	    949048 ns/op
PASS
ok  	github.com/miniohq/logbenchmarks	2.664s

05:34:39 praveen@piper logbenchmarks ±|main ✗|→ go test -bench=. -benchtime=10000x ./...
goos: linux
goarch: amd64
pkg: github.com/miniohq/logbenchmarks
cpu: Intel(R) Core(TM) Ultra 7 165U
BenchmarkJSONSerialization-14             	   10000	       535.2 ns/op
BenchmarkJSONDeserialization-14           	   10000	      2401 ns/op
BenchmarkJSONZstdCompression-14           	   10000	    941646 ns/op
BenchmarkProtobufSerialization-14         	   10000	       283.7 ns/op
BenchmarkProtobufDeserialization-14       	   10000	       344.6 ns/op
BenchmarkProtobufZstdCompression-14       	   10000	    845112 ns/op
BenchmarkMessagePackSerialization-14      	   10000	       577.3 ns/op
BenchmarkMessagePackZstdCompression-14    	   10000	    892619 ns/op
```
