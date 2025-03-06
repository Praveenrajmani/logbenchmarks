

```sh
09:32:48 praveen@piper logbenchmarks ±|main ✗|→ go run .
Size Comparison (Bytes):
JSON: 1980836 bytes
Protobuf: 1581501 bytes
MessagePack: 1766934 bytes
MessagePack Tuple: 1601920 bytes
JSON + Zstd: 102766 bytes
Protobuf + Zstd: 137225 bytes
MessagePack + Zstd: 151922 bytes
MessagePack Tuple + Zstd: 143794 bytes

09:34:33 praveen@piper logbenchmarks ±|main ✗|→ go test -bench=. -benchtime=1000x ./...  -args -batch=10
goos: linux
goarch: amd64
pkg: github.com/miniohq/logbenchmarks
cpu: Intel(R) Core(TM) Ultra 7 165U
BenchmarkJSONSerialization-14                  	    1000	     15242 ns/op	    4705 B/op	      42 allocs/op
BenchmarkJSONDeserialization-14                	    1000	     46763 ns/op	    5681 B/op	      98 allocs/op
BenchmarkJSONZstdCompression-14                	    1000	     34934 ns/op	   24761 B/op	       0 allocs/op
BenchmarkProtobufSerialization-14              	    1000	     14946 ns/op	    2604 B/op	      76 allocs/op
BenchmarkProtobufDeserialization-14            	    1000	     23069 ns/op	    5864 B/op	     165 allocs/op
BenchmarkProtobufZstdCompression-14            	    1000	     35888 ns/op	   24481 B/op	       0 allocs/op
BenchmarkMessagePackSerialization-14           	    1000	      4219 ns/op	    1602 B/op	       0 allocs/op
BenchmarkMessagePackDeserialization-14         	    1000	      6284 ns/op	    1613 B/op	      62 allocs/op
BenchmarkMessagePackZstdCompression-14         	    1000	     31287 ns/op	   24558 B/op	       0 allocs/op
BenchmarkMessagePackTupleSerialization-14      	    1000	      4205 ns/op	    1600 B/op	       0 allocs/op
BenchmarkMessagePackTupleZstdCompression-14    	    1000	     28640 ns/op	   24481 B/op	       0 allocs/op
BenchmarkMessagePackTupleDeserialization-14    	    1000	      6930 ns/op	    1708 B/op	      64 allocs/op
PASS
ok  	github.com/miniohq/logbenchmarks	0.802s

09:34:25 praveen@piper logbenchmarks ±|main ✗|→ go test -bench=. -benchtime=1000x ./...  -args -batch=100
goos: linux
goarch: amd64
pkg: github.com/miniohq/logbenchmarks
cpu: Intel(R) Core(TM) Ultra 7 165U
BenchmarkJSONSerialization-14                  	    1000	     17884 ns/op	    5120 B/op	      45 allocs/op
BenchmarkJSONDeserialization-14                	    1000	     45628 ns/op	    5679 B/op	      98 allocs/op
BenchmarkJSONZstdCompression-14                	    1000	     43306 ns/op	  191383 B/op	       0 allocs/op
BenchmarkProtobufSerialization-14              	    1000	     17470 ns/op	    2719 B/op	      81 allocs/op
BenchmarkProtobufDeserialization-14            	    1000	     23080 ns/op	    5863 B/op	     165 allocs/op
BenchmarkProtobufZstdCompression-14            	    1000	    151314 ns/op	  191548 B/op	       0 allocs/op
BenchmarkMessagePackSerialization-14           	    1000	      3835 ns/op	    1501 B/op	       0 allocs/op
BenchmarkMessagePackDeserialization-14         	    1000	      5818 ns/op	    1614 B/op	      62 allocs/op
BenchmarkMessagePackZstdCompression-14         	    1000	    113028 ns/op	  191712 B/op	       0 allocs/op
BenchmarkMessagePackTupleSerialization-14      	    1000	      3855 ns/op	    1477 B/op	       0 allocs/op
BenchmarkMessagePackTupleZstdCompression-14    	    1000	     94332 ns/op	  191548 B/op	       0 allocs/op
BenchmarkMessagePackTupleDeserialization-14    	    1000	      7928 ns/op	    1708 B/op	      64 allocs/op
PASS
ok  	github.com/miniohq/logbenchmarks	1.037s

09:34:36 praveen@piper logbenchmarks ±|main ✗|→ go test -bench=. -benchtime=1000x ./...  -args -batch=300
goos: linux
goarch: amd64
pkg: github.com/miniohq/logbenchmarks
cpu: Intel(R) Core(TM) Ultra 7 165U
BenchmarkJSONSerialization-14                  	    1000	     15401 ns/op	    6048 B/op	      45 allocs/op
BenchmarkJSONDeserialization-14                	    1000	     44905 ns/op	    5682 B/op	      98 allocs/op
BenchmarkJSONZstdCompression-14                	    1000	     29259 ns/op	   88693 B/op	       0 allocs/op
BenchmarkProtobufSerialization-14              	    1000	     14583 ns/op	    2732 B/op	      81 allocs/op
BenchmarkProtobufDeserialization-14            	    1000	     23012 ns/op	    5853 B/op	     165 allocs/op
BenchmarkProtobufZstdCompression-14            	    1000	     82785 ns/op	   88538 B/op	       0 allocs/op
BenchmarkMessagePackSerialization-14           	    1000	      4635 ns/op	    2250 B/op	       0 allocs/op
BenchmarkMessagePackDeserialization-14         	    1000	      6370 ns/op	    1613 B/op	      62 allocs/op
BenchmarkMessagePackZstdCompression-14         	    1000	     73372 ns/op	   88702 B/op	       0 allocs/op
BenchmarkMessagePackTupleSerialization-14      	    1000	      3349 ns/op	    2193 B/op	       0 allocs/op
BenchmarkMessagePackTupleZstdCompression-14    	    1000	     58176 ns/op	   88563 B/op	       0 allocs/op
BenchmarkMessagePackTupleDeserialization-14    	    1000	      7195 ns/op	    1715 B/op	      64 allocs/op
PASS
ok  	github.com/miniohq/logbenchmarks	0.860s

09:34:46 praveen@piper logbenchmarks ±|main ✗|→ go test -bench=. -benchtime=1000x ./...  -args -batch=500
goos: linux
goarch: amd64
pkg: github.com/miniohq/logbenchmarks
cpu: Intel(R) Core(TM) Ultra 7 165U
BenchmarkJSONSerialization-14                  	    1000	     18577 ns/op	    6347 B/op	      44 allocs/op
BenchmarkJSONDeserialization-14                	    1000	     43356 ns/op	    5674 B/op	      98 allocs/op
BenchmarkJSONZstdCompression-14                	    1000	     46859 ns/op	   54395 B/op	       0 allocs/op
BenchmarkProtobufSerialization-14              	    1000	     18577 ns/op	    2651 B/op	      79 allocs/op
BenchmarkProtobufDeserialization-14            	    1000	     25608 ns/op	    5854 B/op	     165 allocs/op
BenchmarkProtobufZstdCompression-14            	    1000	     52791 ns/op	   54166 B/op	       0 allocs/op
BenchmarkMessagePackSerialization-14           	    1000	      4976 ns/op	    1964 B/op	       0 allocs/op
BenchmarkMessagePackDeserialization-14         	    1000	      6783 ns/op	    1618 B/op	      62 allocs/op
BenchmarkMessagePackZstdCompression-14         	    1000	     47791 ns/op	   54444 B/op	       0 allocs/op
BenchmarkMessagePackTupleSerialization-14      	    1000	      4902 ns/op	    1866 B/op	       0 allocs/op
BenchmarkMessagePackTupleZstdCompression-14    	    1000	     44924 ns/op	   54313 B/op	       0 allocs/op
BenchmarkMessagePackTupleDeserialization-14    	    1000	      5847 ns/op	    1704 B/op	      64 allocs/op
PASS
ok  	github.com/miniohq/logbenchmarks	0.856s

09:34:48 praveen@piper logbenchmarks ±|main ✗|→ go test -bench=. -benchtime=1000x ./...  -args -batch=1000
goos: linux
goarch: amd64
pkg: github.com/miniohq/logbenchmarks
cpu: Intel(R) Core(TM) Ultra 7 165U
BenchmarkJSONSerialization-14                  	    1000	     24924 ns/op	    8911 B/op	      43 allocs/op
BenchmarkJSONDeserialization-14                	    1000	     45686 ns/op	    5678 B/op	      98 allocs/op
BenchmarkJSONZstdCompression-14                	    1000	     28564 ns/op	   35907 B/op	       0 allocs/op
BenchmarkProtobufSerialization-14              	    1000	     17221 ns/op	    2611 B/op	      78 allocs/op
BenchmarkProtobufDeserialization-14            	    1000	     23962 ns/op	    5855 B/op	     165 allocs/op
BenchmarkProtobufZstdCompression-14            	    1000	     32744 ns/op	   36076 B/op	       0 allocs/op
BenchmarkMessagePackSerialization-14           	    1000	      6548 ns/op	    3382 B/op	       0 allocs/op
BenchmarkMessagePackDeserialization-14         	    1000	      6695 ns/op	    1617 B/op	      62 allocs/op
BenchmarkMessagePackZstdCompression-14         	    1000	     34889 ns/op	   36314 B/op	       0 allocs/op
BenchmarkMessagePackTupleSerialization-14      	    1000	      5282 ns/op	    3186 B/op	       0 allocs/op
BenchmarkMessagePackTupleZstdCompression-14    	    1000	     39737 ns/op	   36142 B/op	       0 allocs/op
BenchmarkMessagePackTupleDeserialization-14    	    1000	      6085 ns/op	    1704 B/op	      64 allocs/op
PASS
ok  	github.com/miniohq/logbenchmarks	0.819s
```
