# MARTINI + GORP WITH CONCURRENCY 10

##Apache benchmark

### products#show
Server Hostname:        localhost
Server Port:            3000

Document Path:          /products/1
Document Length:        125 bytes

Concurrency Level:      10
Time taken for tests:   1.149 seconds
Complete requests:      1000
Failed requests:        0
Write errors:           0
Total transferred:      249000 bytes
HTML transferred:       125000 bytes
Requests per second:    870.62 [#/sec] (mean)
Time per request:       11.486 [ms] (mean)
Time per request:       1.149 [ms] (mean, across all concurrent requests)
Transfer rate:          211.70 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.1      0      10
Processing:     1   11   8.0      8      53
Waiting:        0   10   8.1      8      53
Total:          1   11   7.9      9      53

Percentage of the requests served within a certain time (ms)
  50%      9
  66%     14
  75%     17
  80%     18
  90%     22
  95%     24
  98%     31
  99%     35
 100%     53 (longest request)

### products#index
Server Hostname:        localhost
Server Port:            3000

Document Path:          /products?limit=1000
Document Length:        141034 bytes

Concurrency Level:      10
Time taken for tests:   11.027 seconds
Complete requests:      1000
Failed requests:        0
Write errors:           0
Total transferred:      141137000 bytes
HTML transferred:       141034000 bytes
Requests per second:    90.69 [#/sec] (mean)
Time per request:       110.272 [ms] (mean)
Time per request:       11.027 [ms] (mean, across all concurrent requests)
Transfer rate:          12499.04 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.5      0       7
Processing:    20  110  34.4    106     238
Waiting:       20  109  34.3    106     238
Total:         20  110  34.3    107     238

Percentage of the requests served within a certain time (ms)
  50%    107
  66%    120
  75%    129
  80%    135
  90%    155
  95%    172
  98%    194
  99%    214
 100%    238 (longest request)

### products#create
Server Hostname:        localhost
Server Port:            3000

Document Path:          /products
Document Length:        153 bytes

Concurrency Level:      10
Time taken for tests:   1.122 seconds
Complete requests:      1000
Failed requests:        100
   (Connect: 0, Receive: 0, Length: 100, Exceptions: 0)
Write errors:           0
Total transferred:      282072 bytes
Total POSTed:           216216
HTML transferred:       152943 bytes
Requests per second:    891.43 [#/sec] (mean)
Time per request:       11.218 [ms] (mean)
Time per request:       1.122 [ms] (mean, across all concurrent requests)
Transfer rate:          245.55 [Kbytes/sec] received
                        188.22 kb/s sent
                        433.78 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.5      0       4
Processing:     2   11   6.1      8      30
Waiting:        0   11   6.1      8      29
Total:          2   11   6.0      9      30

Percentage of the requests served within a certain time (ms)
  50%      9
  66%     15
  75%     17
  80%     18
  90%     20
  95%     21
  98%     23
  99%     26
 100%     30 (longest request)

### products#update
Server Hostname:        localhost
Server Port:            3000

Document Path:          /products/1
Document Length:        133 bytes

Concurrency Level:      10
Time taken for tests:   1.201 seconds
Complete requests:      1000
Failed requests:        113
   (Connect: 0, Receive: 0, Length: 113, Exceptions: 0)
Write errors:           0
Total transferred:      256874 bytes
Total PUT:              217000
HTML transferred:       132874 bytes
Requests per second:    832.67 [#/sec] (mean)
Time per request:       12.010 [ms] (mean)
Time per request:       1.201 [ms] (mean, across all concurrent requests)
Transfer rate:          208.88 [Kbytes/sec] received
                        176.46 kb/s sent
                        385.33 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.7      0      11
Processing:     2   11   7.0      9      47
Waiting:        0   11   7.0      8      47
Total:          2   12   6.9      9      47

Percentage of the requests served within a certain time (ms)
  50%      9
  66%     15
  75%     18
  80%     19
  90%     22
  95%     24
  98%     28
  99%     30
 100%     47 (longest request)

### products#bulk_create
Server Hostname:        localhost
Server Port:            3000

Document Path:          /products/bulk
Document Length:        12231 bytes

Concurrency Level:      10
Time taken for tests:   23.619 seconds
Complete requests:      1000
Failed requests:        0
Write errors:           0
Total transferred:      12339000 bytes
Total POSTed:           8613000
HTML transferred:       12231000 bytes
Requests per second:    42.34 [#/sec] (mean)
Time per request:       236.194 [ms] (mean)
Time per request:       23.619 [ms] (mean, across all concurrent requests)
Transfer rate:          510.16 [Kbytes/sec] received
                        356.11 kb/s sent
                        866.28 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:    95  235  55.4    227     559
Waiting:       95  235  55.4    227     559
Total:         95  236  55.4    227     559

Percentage of the requests served within a certain time (ms)
  50%    227
  66%    251
  75%    268
  80%    280
  90%    308
  95%    337
  98%    371
  99%    389
 100%    559 (longest request)


## Siege

### products#index

GET /products?limit=100

Transactions:		      100000 hits
Availability:		      100.00 %
Elapsed time:		      483.20 secs
Data transferred:	     1333.71 MB
Response time:		        0.46 secs
Transaction rate:	      206.95 trans/sec
Throughput:		        2.76 MB/sec
Concurrency:		       95.50
Successful transactions:      100000
Failed transactions:	           0
Longest transaction:	       15.65
Shortest transaction:	        0.00
