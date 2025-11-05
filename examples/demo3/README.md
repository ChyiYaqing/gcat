# Set up

```bash
gcat/examples/demo3 on  main [!+?] via 🐹 v1.25.3 
➜ go test -bench=BenchmarkUnary$ .
goos: darwin
goarch: arm64
pkg: github.com/chyiyaqing/gcat/examples/demo3
cpu: Apple M2
BenchmarkUnary-8           24160             47069 ns/op
PASS
ok      github.com/chyiyaqing/gcat/examples/demo3       2.994s

gcat/examples/demo3 on  main [!+?] via 🐹 v1.25.3 took 3.7s 
➜ VTCODEC=1 go test -bench=BenchmarkUnary$ .
goos: darwin
goarch: arm64
pkg: github.com/chyiyaqing/gcat/examples/demo3
cpu: Apple M2
BenchmarkUnary-8           24076             48767 ns/op
PASS
ok      github.com/chyiyaqing/gcat/examples/demo3       1.869s
```

性能分析

```bash

gcat/examples/demo3 on  main [!+?] via 🐹 v1.25.3 took 27.5s 
➜ go test -bench=BenchmarkMarshalProto -benchmem -count=10 | tee  marshal_default.txt

goos: darwin
goarch: arm64
pkg: github.com/chyiyaqing/gcat/examples/demo3
cpu: Apple M2
BenchmarkMarshalProto-8          7853743               148.6 ns/op            64 B/op          1 allocs/op
BenchmarkMarshalProto-8          8809044               136.2 ns/op            64 B/op          1 allocs/op
BenchmarkMarshalProto-8          8687488               136.5 ns/op            64 B/op          1 allocs/op
BenchmarkMarshalProto-8          8760613               136.1 ns/op            64 B/op          1 allocs/op
BenchmarkMarshalProto-8          8708493               135.9 ns/op            64 B/op          1 allocs/op
BenchmarkMarshalProto-8          8716977               136.2 ns/op            64 B/op          1 allocs/op
BenchmarkMarshalProto-8          8628021               138.2 ns/op            64 B/op          1 allocs/op
BenchmarkMarshalProto-8          8709829               139.7 ns/op            64 B/op          1 allocs/op
BenchmarkMarshalProto-8          8699161               136.1 ns/op            64 B/op          1 allocs/op
BenchmarkMarshalProto-8          8696956               135.8 ns/op            64 B/op          1 allocs/op
PASS
ok      github.com/chyiyaqing/gcat/examples/demo3       14.483s

gcat/examples/demo3 on  main [!+?] via 🐹 v1.25.3 took 15.6s 
➜ go test -bench=BenchmarkMarshalVT -benchmem -count=10 | tee  marshal_vt.txt        

goos: darwin
goarch: arm64
pkg: github.com/chyiyaqing/gcat/examples/demo3
cpu: Apple M2
BenchmarkMarshalVT-8    31061119                38.86 ns/op           64 B/op          1 allocs/op
BenchmarkMarshalVT-8    29827158                39.14 ns/op           64 B/op          1 allocs/op
BenchmarkMarshalVT-8    29468733                39.04 ns/op           64 B/op          1 allocs/op
BenchmarkMarshalVT-8    29949895                38.96 ns/op           64 B/op          1 allocs/op
BenchmarkMarshalVT-8    29945598                39.14 ns/op           64 B/op          1 allocs/op
BenchmarkMarshalVT-8    29924036                39.11 ns/op           64 B/op          1 allocs/op
BenchmarkMarshalVT-8    29455412                39.20 ns/op           64 B/op          1 allocs/op
BenchmarkMarshalVT-8    29900145                39.35 ns/op           64 B/op          1 allocs/op
BenchmarkMarshalVT-8    29803575                39.05 ns/op           64 B/op          1 allocs/op
BenchmarkMarshalVT-8    30087189                39.39 ns/op           64 B/op          1 allocs/op
PASS
ok      github.com/chyiyaqing/gcat/examples/demo3       14.129s

gcat/examples/demo3 on  main [!+?] via 🐹 v1.25.3 took 15.2s 
➜ benchstat marshal_default.txt marshal_vt.txt
goos: darwin
goarch: arm64
pkg: github.com/chyiyaqing/gcat/examples/demo3
cpu: Apple M2
               │ marshal_default.txt │   marshal_vt.txt   │
               │       sec/op        │   sec/op     vs base   │
MarshalProto-8      136.2n ± 3%      | MarshalVT-8 39.12n ± 1%
geomean                  136.2n        39.12n       ? ¹ ²
¹ benchmark set differs from baseline; geomeans may not be comparable
² ratios must be >0 to compute geomean

               │ marshal_default.txt │  marshal_vt.txt   │
               │        B/op         │    B/op     vs base   │
MarshalProto-8            64.00 ± 0%
MarshalVT-8                            64.00 ± 0%
geomean                   64.00        64.00       ? ¹ ²
¹ benchmark set differs from baseline; geomeans may not be comparable
² ratios must be >0 to compute geomean

               │ marshal_default.txt │  marshal_vt.txt   │
               │      allocs/op      │ allocs/op   vs base   │
MarshalProto-8            1.000 ± 0%
MarshalVT-8                            1.000 ± 0%
geomean                   1.000        1.000       ? ¹ ²
¹ benchmark set differs from baseline; geomeans may not be comparable
² ratios must be >0 to compute geomean

gcat/examples/demo3 on  main [!+?] via 🐹 v1.25.3 
➜ 
```