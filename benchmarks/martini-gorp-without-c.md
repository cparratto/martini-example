# MARTINI + GORP WITHOUT CONCURRENCY

## products#show
Server Hostname:        localhost
Server Port:            3000

Document Path:          /products/1
Document Length:        136 bytes

Concurrency Level:      1
Time taken for tests:   0.661 seconds
Complete requests:      1000
Failed requests:        0
Write errors:           0
Total transferred:      260000 bytes
HTML transferred:       136000 bytes
Requests per second:    1513.53 [#/sec] (mean)
Time per request:       0.661 [ms] (mean)
Time per request:       0.661 [ms] (mean, across all concurrent requests)
Transfer rate:          384.30 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     0    1   0.1      1       4
Waiting:        0    1   0.1      1       4
Total:          1    1   0.1      1       4

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      1
  95%      1
  98%      1
  99%      1
 100%      4 (longest request)

## products#index
Server Hostname:        localhost
Server Port:            3000

Document Path:          /products?limit=1000
Document Length:        141029 bytes

Concurrency Level:      1
Time taken for tests:   10.009 seconds
Complete requests:      1000
Failed requests:        0
Write errors:           0
Total transferred:      141132000 bytes
HTML transferred:       141029000 bytes
Requests per second:    99.91 [#/sec] (mean)
Time per request:       10.009 [ms] (mean)
Time per request:       10.009 [ms] (mean, across all concurrent requests)
Transfer rate:          13769.45 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     8   10   1.2      9      16
Waiting:        8   10   1.2      9      16
Total:          8   10   1.2     10      17

Percentage of the requests served within a certain time (ms)
  50%     10
  66%     10
  75%     11
  80%     11
  90%     12
  95%     12
  98%     13
  99%     13
 100%     17 (longest request)

## products#create
Server Hostname:        localhost
Server Port:            3000

Document Path:          /products
Document Length:        153 bytes

Concurrency Level:      1
Time taken for tests:   0.960 seconds
Complete requests:      1000
Failed requests:        95
   (Connect: 0, Receive: 0, Length: 95, Exceptions: 0)
Write errors:           0
Total transferred:      281792 bytes
Total POSTed:           216000
HTML transferred:       152792 bytes
Requests per second:    1041.42 [#/sec] (mean)
Time per request:       0.960 [ms] (mean)
Time per request:       0.960 [ms] (mean, across all concurrent requests)
Transfer rate:          286.59 [Kbytes/sec] received
                        219.68 kb/s sent
                        506.26 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     1    1   0.3      1       7
Waiting:        1    1   0.3      1       7
Total:          1    1   0.3      1       7

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      1
  95%      1
  98%      1
  99%      2
 100%      7 (longest request)

## products#update
Server Hostname:        localhost
Server Port:            3000

Document Path:          /products/1
Document Length:        133 bytes

Concurrency Level:      1
Time taken for tests:   0.983 seconds
Complete requests:      1000
Failed requests:        101
   (Connect: 0, Receive: 0, Length: 101, Exceptions: 0)
Write errors:           0
Total transferred:      256879 bytes
Total PUT:              217000
HTML transferred:       132879 bytes
Requests per second:    1017.72 [#/sec] (mean)
Time per request:       0.983 [ms] (mean)
Time per request:       0.983 [ms] (mean, across all concurrent requests)
Transfer rate:          255.30 [Kbytes/sec] received
                        215.67 kb/s sent
                        470.97 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     1    1   0.2      1       5
Waiting:        1    1   0.2      1       5
Total:          1    1   0.2      1       5

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      1
  95%      1
  98%      1
  99%      2
 100%      5 (longest request)

## products#bulk_create
Server Hostname:        localhost
Server Port:            3000

Document Path:          /products/bulk
Document Length:        12231 bytes

Concurrency Level:      1
Time taken for tests:   67.158 seconds
Complete requests:      1000
Failed requests:        0
Write errors:           0
Total transferred:      12339000 bytes
Total POSTed:           8613000
HTML transferred:       12231000 bytes
Requests per second:    14.89 [#/sec] (mean)
Time per request:       67.158 [ms] (mean)
Time per request:       67.158 [ms] (mean, across all concurrent requests)
Transfer rate:          179.42 [Kbytes/sec] received
                        125.24 kb/s sent
                        304.67 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:    52   67  12.1     71     149
Waiting:       52   67  12.1     71     149
Total:         53   67  12.2     71     149

Percentage of the requests served within a certain time (ms)
  50%     71
  66%     75
  75%     77
  80%     78
  90%     79
  95%     80
  98%     93
  99%    101
 100%    149 (longest request)


## Siege

### products#index

GET /products?limit=100

Transactions:		        1000 hits
Availability:		      100.00 %
Elapsed time:		        2.73 secs
Data transferred:	       13.34 MB
Response time:		        0.00 secs
Transaction rate:	      366.30 trans/sec
Throughput:		        4.89 MB/sec
Concurrency:		        0.98
Successful transactions:        1000
Failed transactions:	           0
Longest transaction:	        0.03
Shortest transaction:	        0.00
