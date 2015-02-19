# Benchmark

- net.go

ab -n 1000 -c 100 http://10.211.55.14:9090/rankup?bonus=fGhudaLda

```
Server Software:
Server Hostname:        10.211.55.14
Server Port:            9090

Document Path:          /rankup?bonus=fGhudaLda
Document Length:        19 bytes

Concurrency Level:      100
Time taken for tests:   0.073 seconds
Complete requests:      1000
Failed requests:        0
Write errors:           0
Non-2xx responses:      1000
Total transferred:      143000 bytes
HTML transferred:       19000 bytes
Requests per second:    13742.12 [#/sec] (mean)
Time per request:       7.277 [ms] (mean)
Time per request:       0.073 [ms] (mean, across all concurrent requests)
Transfer rate:          1919.06 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.7      0       3
Processing:     1    6   1.9      6      12
Waiting:        1    6   1.9      6      11
Total:          3    7   1.5      7      12
WARNING: The median and mean for the initial connection time are not within a normal deviation
```

- rankup.go

```
Server Software:
Server Hostname:        10.211.55.14
Server Port:            9090

Document Path:          /rankup?bonus=fGhudaLda
Document Length:        35 bytes

Concurrency Level:      100
Time taken for tests:   0.200 seconds
Complete requests:      1000
Failed requests:        0
Write errors:           0
Total transferred:      152000 bytes
HTML transferred:       35000 bytes
Requests per second:    5008.89 [#/sec] (mean)
Time per request:       19.964 [ms] (mean)
Time per request:       0.200 [ms] (mean, across all concurrent requests)
Transfer rate:          743.51 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.8      0       3
Processing:     1   19   6.3     19      36
Waiting:        1   19   6.3     19      36
Total:          3   19   5.9     19      36
```

- net.go(build)

```

Server Software:
Server Hostname:        10.211.55.14
Server Port:            9090

Document Path:          /rankup?bonus=fGhudaLda
Document Length:        19 bytes

Concurrency Level:      100
Time taken for tests:   0.087 seconds
Complete requests:      1000
Failed requests:        0
Write errors:           0
Non-2xx responses:      1001
Total transferred:      143143 bytes
HTML transferred:       19019 bytes
Requests per second:    11489.50 [#/sec] (mean)
Time per request:       8.704 [ms] (mean)
Time per request:       0.087 [ms] (mean, across all concurrent requests)
Transfer rate:          1606.10 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.3      0      10
Processing:     1    7   4.1      7      20
Waiting:        0    7   2.6      7      16
Total:          3    8   3.7      7      20

```

- rankup.go(build)

```

Server Software:
Server Hostname:        10.211.55.14
Server Port:            9090

Document Path:          /rankup?bonus=fGhudaLda
Document Length:        35 bytes

Concurrency Level:      100
Time taken for tests:   0.189 seconds
Complete requests:      1000
Failed requests:        0
Write errors:           0
Total transferred:      152000 bytes
HTML transferred:       35000 bytes
Requests per second:    5284.32 [#/sec] (mean)
Time per request:       18.924 [ms] (mean)
Time per request:       0.189 [ms] (mean, across all concurrent requests)
Transfer rate:          784.39 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.7      0       3
Processing:     5   18   6.1     18      35
Waiting:        5   18   6.0     18      33
Total:          5   18   5.8     18      35
```
