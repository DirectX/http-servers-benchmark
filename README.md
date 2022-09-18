# HTTP Servers Benchmark

## Tests
* nodejs-express
* go-http
* rust-nickel

## Running

### Node JS

```bash
$ cd nodejs-...
$ npm start
```

### Node JS

```bash
$ cd go-...
$ go run .
```

### Rust

```bash
$ cd rust-...
$ cargo run
```

## Results

Measured at Ryzen 5950x (16 cores / 32 threads), 128 GB RAM, Samsung EVO 980 PRO SSD 1TB, localhost

### Node JS + Express

```bash
$ wrk -t12 -c400 -d30s http://127.0.0.1:3000/
Running 30s test @ http://127.0.0.1:3000/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    33.18ms    5.34ms 165.66ms   94.81%
    Req/Sec     1.01k    81.45     1.24k    81.83%
  359349 requests in 30.02s, 81.91MB read
Requests/sec:  11970.16
Transfer/sec:      2.73MB
```

### Go + HTTP
```bash
$ wrk -t12 -c400 -d30s http://127.0.0.1:4000/
Running 30s test @ http://127.0.0.1:4000/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   693.11us    1.85ms  97.74ms   96.53%
    Req/Sec    66.58k     5.67k  132.96k    74.22%
  23884657 requests in 30.10s, 2.87GB read
Requests/sec: 793537.20
Transfer/sec:     97.62MB
```

### Rust + Nickel
```bash
$ wrk -t12 -c400 -d30s http://127.0.0.1:5000/
Running 30s test @ http://127.0.0.1:5000/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   164.50us   97.94us  22.53ms   91.63%
    Req/Sec    56.80k    31.57k  116.33k    49.21%
  6788093 requests in 30.10s, 0.93GB read
Requests/sec: 225523.47
Transfer/sec:     31.62MB
```