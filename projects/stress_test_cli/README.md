# Stress test CLI

## Objective

Create a CLI in Golang to stress test web services. The user provides URL, total number of requests and number of concurrent operations.

```sh
Flags

--url: "Web servers url address"
--requests: "Total number of requests to be sent"
--concurrency: "Number of concurrent operations"

```

## Test flow

Build image:

```
docker build -t stress_test_cli:latest .

```

Run CLI in container:

```

docker run stress_test_cli:latest --url=https://www.google.com/ --requests=100 --concurrency=10

```

Example result:

```
====================================
    STRESS TEST REPORT
====================================
Target URL:           https://www.google.com/
Total Requests:       100
Concurrency Level:    10
Total Elapsed Time:   14.336s

Results:
Successful (200):  100
Failed:            0

====================================

```
